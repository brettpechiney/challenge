package config

import (
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/brettpechiney/challenge/config/param"
)

// Challenge is is a configuration implementation backed by Viper.
type Challenge struct {
	remote bool
	v      *viper.Viper
}

// Load returns a Challenge config that loads configuration properties
// from a file. The file is kept at project root with CI/CD in mind.
func Load(configPaths []string) (*Challenge, error) {
	i := &Challenge{v: viper.New()}
	i.setDefaults()
	i.setupEnvVarReader()

	for _, dir := range configPaths {
		i.v.AddConfigPath(dir)
	}

	i.v.SetConfigName("application-properties")
	i.v.SetConfigType("toml")

	if err := i.v.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrapf(err, "unable to read configuration file")
		}
		log.Printf("no configuration file found; proceeding without one")
	}

	return i, nil
}

// Defaults returns a Challenge that has just the default values
// set. It will load neither local nor remote files.
func Defaults() *Challenge {
	i := &Challenge{v: viper.New()}
	i.setDefaults()
	return i
}

// Set overrides the configuration value. It is used for testing.
func (i *Challenge) Set(key string, value interface{}) {
	i.v.Set(key, value)
}

// DatabasePrefix returns the prefix of the database that stores Challenge
// application information.
func (i *Challenge) DatabasePrefix() string {
	return i.v.GetString(param.DatabasePrefix)
}

// DatabaseUser returns the database user associated with the
// credentials.
func (i *Challenge) DatabaseUser() string {
	return i.v.GetString(param.DatabaseUser)
}

// DatabasePassword returns the database credentials.
func (i *Challenge) DatabasePassword() string {
	return i.v.GetString(param.DatabasePassword)
}

// DatabaseHost returns the host the database is running on.
func (i *Challenge) DatabaseHost() string {
	return i.v.GetString(param.DatabaseHost)
}

// DatabasePort returns the port the database is listening on.
func (i *Challenge) DatabasePort() string {
	return i.v.GetString(param.DatabasePort)
}

// DatabaseName returns the name of the database.
func (i *Challenge) DatabaseName() string {
	return i.v.GetString(param.DatabaseName)
}

// LoggingLevel returns the application's logging level.
func (i *Challenge) LoggingLevel() string {
	return i.v.GetString(param.LoggingLevel)
}

func (i *Challenge) setDefaults() {
	i.v.SetDefault(param.DatabasePrefix, "postgresql://")
	i.v.SetDefault(param.DatabaseUser, "maxroach")
	i.v.SetDefault(param.DatabasePassword, "maxroach")
	i.v.SetDefault(param.DatabaseHost, "localhost")
	i.v.SetDefault(param.DatabasePort, 26257)
	i.v.SetDefault(param.DatabaseName, "challenge")

	i.v.SetDefault(param.LoggingLevel, "INFO")
}

func (i *Challenge) setupEnvVarReader() {
	i.v.AutomaticEnv()
	i.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
