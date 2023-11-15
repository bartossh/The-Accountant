package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"

	"github.com/bartossh/Computantis/accountant"
	"github.com/bartossh/Computantis/aeswrapper"
	"github.com/bartossh/Computantis/configuration"
	"github.com/bartossh/Computantis/dataprovider"
	"github.com/bartossh/Computantis/fileoperations"
	"github.com/bartossh/Computantis/gossip"
	"github.com/bartossh/Computantis/logging"
	"github.com/bartossh/Computantis/logo"
	"github.com/bartossh/Computantis/natsclient"
	"github.com/bartossh/Computantis/notaryserver"
	"github.com/bartossh/Computantis/stdoutwriter"
	"github.com/bartossh/Computantis/telemetry"
	"github.com/bartossh/Computantis/wallet"
	"github.com/bartossh/Computantis/zincaddapter"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

const usage = `runs the Computantis node that connects in to the Computantis network`

const (
	rxBufferSize  = 100
	vrxBufferSize = 100
)

const gossipTimeout = time.Second * 5

func main() {
	logo.Display()

	var file string
	configurator := func() (configuration.Configuration, error) {
		if file == "" {
			return configuration.Configuration{}, errors.New("please specify configuration file path with -c <path to file>")
		}

		cfg, err := configuration.Read(file)
		if err != nil {
			return cfg, err
		}

		return cfg, nil
	}

	app := &cli.App{
		Name:  "computantis",
		Usage: usage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `FILE`",
				Destination: &file,
			},
		},
		Action: func(_ *cli.Context) error {
			cfg, err := configurator()
			if err != nil {
				return err
			}
			run(cfg)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		pterm.Error.Println(err.Error())
	}
}

func run(cfg configuration.Configuration) {
	if cfg.IsProfiling {
		f, _ := os.Create("default.pgo")
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		cancel()
	}()

	callbackOnErr := func(err error) {
		fmt.Println("Error with logger: ", err)
	}

	callbackOnFatal := func(err error) {
		panic(fmt.Sprintf("Error with logger: %s", err))
	}

	zinc, err := zincaddapter.New(cfg.ZincLogger)
	if err != nil {
		fmt.Println(err)
		c <- os.Interrupt
		return
	}
	log := logging.New(callbackOnErr, callbackOnFatal, stdoutwriter.Logger{}, &zinc)
	dataProvider := dataprovider.New(ctx, cfg.DataProvider)
	verifier := wallet.NewVerifier()
	vrxCh := make(chan *accountant.Vertex)
	h := fileoperations.New(cfg.FileOperator, aeswrapper.New())
	wlt, err := h.ReadWallet()
	if err != nil {
		log.Error(err.Error())
		c <- os.Interrupt
		return
	}

	acc, err := accountant.NewAccountingBook(ctx, cfg.Accountant, &verifier, &wlt, &log)
	if err != nil {
		log.Error(err.Error())
		c <- os.Interrupt
		return
	}

	tele, err := telemetry.Run(ctx, cancel, 2112)
	if err != nil {
		log.Error(err.Error())
		c <- os.Interrupt
		return
	}

	pub, err := natsclient.PublisherConnect(cfg.Nats)
	if err != nil {
		log.Error(err.Error())
		c <- os.Interrupt
		return
	}
	defer func() {
		if err := pub.Disconnect(); err != nil {
			log.Error(err.Error())
		}
	}()

	go func() {
		err = gossip.RunGRPC(ctx, cfg.Gossip, &log, gossipTimeout, &wlt, &verifier, acc, vrxCh)
		if err != nil {
			log.Error(err.Error())
			c <- os.Interrupt
			return
		}
	}()

	err = notaryserver.Run(ctx, cfg.NotaryServer, pub, dataProvider, tele, &log, &verifier, acc, vrxCh)
	if err != nil {
		log.Error(err.Error())
	}
	time.Sleep(time.Second)
}
