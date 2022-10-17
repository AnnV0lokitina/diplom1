package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

// DB keep information about database connection.
type DB struct {
	conn *pgx.Conn
}

// NewDB Creates a new database connection.
func NewDB(ctx context.Context, dsn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

// Close the database connection.
func (r *DB) Close(ctx context.Context) error {
	return r.conn.Close(ctx)
}
