package repohelper

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/bartossh/Computantis/block"
	"github.com/bartossh/Computantis/repomongo"
	"github.com/bartossh/Computantis/repopostgre"
	"github.com/bartossh/Computantis/transaction"
	"github.com/bartossh/Computantis/validator"
	"github.com/lib/pq"
)

var (
	ErrDatabaseNotSupported = fmt.Errorf("database not supported")
)

// AddressWriteFindChecker abstracts address operations.
type AddressWriteFindChecker interface {
	WriteAddress(ctx context.Context, addr string) error
	CheckAddressExists(ctx context.Context, addr string) (bool, error)
	FindAddress(ctx context.Context, search string, limit int) ([]string, error)
}

// BlockReadWriter abstracts block operations.
type BlockReadWriter interface {
	LastBlock(ctx context.Context) (block.Block, error)
	ReadBlockByHash(ctx context.Context, hash [32]byte) (block.Block, error)
	WriteBlock(ctx context.Context, block block.Block) error
}

// MigrationRunner abstracts migration operations.
type Migrator interface {
	RunMigration(ctx context.Context) error
}

// TokenWriteCheckInvalidator abstracts token operations.
type TokenWriteCheckInvalidator interface {
	CheckToken(ctx context.Context, tkn string) (bool, error)
	WriteToken(ctx context.Context, tkn string, expirationDate int64) error
	InvalidateToken(ctx context.Context, token string) error
}

// TransactionOperator abstracts transaction operations.
type TransactionOperator interface {
	WriteTransactionsInBlock(ctx context.Context, blockHash [32]byte, trxHash [][32]byte) error
	FindTransactionInBlockHash(ctx context.Context, trxHash [32]byte) ([32]byte, error)
	WriteTemporaryTransaction(ctx context.Context, trx *transaction.Transaction) error
	RemoveAwaitingTransaction(ctx context.Context, trxHash [32]byte) error
	WriteIssuerSignedTransactionForReceiver(ctx context.Context, receiverAddr string, trx *transaction.Transaction) error
	ReadAwaitingTransactionsByReceiver(ctx context.Context, address string) ([]transaction.Transaction, error)
	ReadAwaitingTransactionsByIssuer(ctx context.Context, address string) ([]transaction.Transaction, error)
	MoveTransactionsFromTemporaryToPermanent(ctx context.Context, hash [][32]byte) error
	ReadTemporaryTransactions(ctx context.Context) ([]transaction.Transaction, error)
}

// ValidatorStatusReader abstracts validator status operations.
type ValidatorStatusReader interface {
	ReadLastNValidatorStatuses(ctx context.Context, last int64) ([]validator.Status, error)
	WriteValidatorStatus(ctx context.Context, vs *validator.Status) error
}

// Subscriber abstracts blockchain subscription to blockchain locks.
type Subscriber interface {
	SubscribeToLockBlockchainNotification(ctx context.Context, c chan<- bool, node string)
}

// Synchronizer abstracts blockchain synchronization operations.
type Synchronizer interface {
	AddToBlockchainLockQueue(ctx context.Context, nodeID string) error
	RemoveFromBlockchainLocks(ctx context.Context, nodeID string) error
	CheckIsOnTopOfBlockchainsLocks(ctx context.Context, nodeID string) (bool, error)
}

// NodeRegister abstracts node registration operations.
type NodeRegister interface {
	RegisterNode(ctx context.Context, n, ws string) error
	UnregisterNode(ctx context.Context, n string) error
	CountRegistered(ctx context.Context) (int, error)
	ReadRegisteredNodesAddresses(ctx context.Context) ([]string, error)
}

// ConnectionCloser abstracts connection closing operations.
type ConnectionCloser interface {
	Disconnect(ctx context.Context) error
}

// RepositoryProvider is an interface that ensures that all required methods to run computantis are implemented.
type RepositoryProvider interface {
	AddressWriteFindChecker
	BlockReadWriter
	io.Writer
	Migrator
	TokenWriteCheckInvalidator
	TransactionOperator
	ValidatorStatusReader
	Synchronizer
	NodeRegister
	ConnectionCloser
}

// Config contains configuration for the database.
type DBConfig struct {
	ConnStr      string `yaml:"conn_str"`         // ConnStr is the connection string to the database.
	DatabaseName string `yaml:"database_name"`    // DatabaseName is the name of the database.
	Token        string `yaml:"token"`            // Token is the token that is used to confirm api clients access.
	TokenExpire  int64  `yaml:"token_expiration"` // TokenExpire is the number of seconds after which token expires.
}

// Connect connects to the proper database and returns that connection.
func (cfg DBConfig) Connect(ctx context.Context) (RepositoryProvider, Subscriber, error) {
	switch {
	case strings.Contains(cfg.ConnStr, "postgres"):
		db, err := repopostgre.Connect(ctx, cfg.ConnStr, cfg.DatabaseName)
		if err != nil {
			return nil, nil, err
		}
		f := func(ev pq.ListenerEventType, err error) {
			if err != nil {
				panic(err)
			}
		}
		lister, err := repopostgre.Listen(cfg.ConnStr, f)
		if err != nil {
			return nil, nil, err
		}
		return db, lister, nil
	case strings.Contains(cfg.ConnStr, "mongodb"):
		db, err := repomongo.Connect(ctx, cfg.ConnStr, cfg.DatabaseName)
		if err != nil {
			return nil, nil, err
		}
		return db, db, nil
	}

	return nil, nil, ErrDatabaseNotSupported
}