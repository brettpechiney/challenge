package config_test

import (
	"fmt"
	"github.com/brettpechiney/challenge/config"
	"testing"
)

// Test getters by comparing against default values.
func TestGetters(t *testing.T) {
	cfg := config.Defaults()
	testCases := []struct {
		Name         string
		TestFunction func() string
		Expected     string
	}{
		{
			"DatabasePrefix",
			cfg.DatabasePrefix,
			"postgresql://",
		},
		{
			"DatabaseUser",
			cfg.DatabaseUser,
			"maxroach",
		},
		{
			"DatabasePassword",
			cfg.DatabasePassword,
			"maxroach",
		},
		{
			"DatabaseHost",
			cfg.DatabaseHost,
			"localhost",
		},
		{
			"DatabasePort",
			cfg.DatabasePort,
			"26257",
		},
		{
			"DatabaseName",
			cfg.DatabaseName,
			"challenge",
		},
		{
			"LoggingLevel",
			cfg.LoggingLevel,
			"INFO",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			if actual := tc.TestFunction(); actual != tc.Expected {
				t.Errorf("expected '%s', got '%s'", tc.Expected, actual)
			}
		})
	}
}
