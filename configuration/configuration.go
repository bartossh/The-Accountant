package configuration

import (
	"fmt"
	"os"

	"github.com/bartossh/The-Accountant/bookkeeping"
	"github.com/bartossh/The-Accountant/dataprovider"
	"github.com/bartossh/The-Accountant/repo"
	"github.com/bartossh/The-Accountant/server"
	"gopkg.in/yaml.v2"
)

// Main is the main configuration of the application that corresponds to the *.yaml file
// that holds the configuration.
type Main struct {
	Bookkeeper   bookkeeping.Config  `yaml:"bookkeeper"`
	Server       server.Config       `yaml:"server"`
	Database     repo.Config         `yaml:"database"`
	DataProvider dataprovider.Config `yaml:"data_provider"`
}

// Read reads the configuration  from the file and returns the MainYaml struct.
func Read(path string) (Main, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return Main{}, err
	}

	var main Main
	err = yaml.Unmarshal(buf, &main)
	if err != nil {
		return Main{}, fmt.Errorf("in file %q: %w", path, err)
	}

	return main, err
}