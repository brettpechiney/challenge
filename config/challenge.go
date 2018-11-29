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

// DataSource returns the connection string of the database that
// stores Challenge application information.
func (i *Challenge) DataSource() string {
	return i.v.GetString(param.DataSource)
}

// LoggingLevel returns the application's logging level.
func (i *Challenge) LoggingLevel() string {
	return i.v.GetString(param.LoggingLevel)
}

// SigningKey returns the application's JWT signing key.
func (i *Challenge) SigningKey() string {
	return i.v.GetString(param.SigningKey)
}

func (i *Challenge) setDefaults() {
	const Source = "postgresql://maxroach@localhost:26257/challenge?sslmode=disable"
	const Level = "INFO"
	const SigningKey = "supersecretkey"
	i.v.SetDefault(param.DataSource, Source)
	i.v.SetDefault(param.LoggingLevel, Level)
	i.v.SetDefault(param.SigningKey, SigningKey)
}

func (i *Challenge) setupEnvVarReader() {
	i.v.AutomaticEnv()
	i.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
