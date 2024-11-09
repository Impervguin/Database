package task6

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
)

const (
	maxConn        = 10
	acquireTimeout = time.Minute
)

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}

func NewTask6Storage(ctx context.Context) (*Task6Storage, error) {
	conf, err := pgx.ParseURI(DSN())

	if err != nil {
		return nil, err
	}
	poolConf := &pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: maxConn,
		AcquireTimeout: acquireTimeout,
	}
	conn, err := pgx.NewConnPool(*poolConf)
	if err != nil {
		return nil, err
	}

	return &Task6Storage{conn: conn}, nil
}

type Task6Storage struct {
	conn *pgx.ConnPool
}
