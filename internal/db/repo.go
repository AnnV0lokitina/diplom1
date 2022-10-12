package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type DB struct {
	conn *pgx.Conn
}

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

func (r *DB) Close(ctx context.Context) error {
	return r.conn.Close(ctx)
}
