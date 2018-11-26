package cockroach

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq" // Import underlying database driver.
	"github.com/pkg/errors"
)

// DAO provides application-level context to the database handle.
type DAO struct {
	*sql.DB
}

// Tx provides application-level context to the transaction handle.
type Tx struct {
	*sql.Tx
}

var ctx = context.Background()

// NewDAO creates a database object, associates it with the
// Postgres driver, and validates the database connection.
func NewDAO(dataSource string) (*DAO, error) {
	conn, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open connection with CockroachDB")
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err = conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, errors.Wrapf(err, "failed to ping CockroachDB")
	}
	return &DAO{conn}, nil
}
