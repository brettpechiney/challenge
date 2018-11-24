package data

import (
	"context"
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq" // Import underlying database driver.
	"github.com/pkg/errors"

	"github.com/brettpechiney/challenge/config"
)

// The DB struct provides application-level context to the
// database handle.
type DB struct {
	*sql.DB
}

// The Tx struct provides application-level context to the
// transaction handle.
type Tx struct {
	*sql.Tx
}

var ctx = context.Background()

// New ceates a database object, associates it with the
// Postgres driver, and validates the database connection.
func New(cfg *config.Challenge) (*DB, error) {
	sourceName := connectionString(cfg)

	conn, err := sql.Open("postgres", sourceName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open connection with CockroachDB")
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err = conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, errors.Wrapf(err, "failed to ping CockroachDB")
	}
	return &DB{conn}, nil
}

func connectionString(cfg *config.Challenge) string {
	// TODO: implement secure connection if you have time.
	var sb strings.Builder
	sb.WriteString(cfg.DatabasePrefix())
	sb.WriteString(cfg.DatabaseUser())
	sb.WriteString("@")
	sb.WriteString(cfg.DatabaseHost())
	sb.WriteString(":")
	sb.WriteString(cfg.DatabasePort())
	sb.WriteString("/")
	sb.WriteString(cfg.DatabaseName())
	sb.WriteString("?sslmode=disable")

	return sb.String()
}
