package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	pgUrl   string
	Conn    *pgx.Conn
	NoCA    int
	Timeout int
	logger  *slog.Logger
}

func NewDB(pgUrl string, noca, tim int, log *slog.Logger) *DB {
	return &DB{
		pgUrl:   pgUrl,
		NoCA:    noca,
		Timeout: tim,
		logger:  log,
	}
}

func (db *DB) ConnectWithDB() (err error) {
	for i := 0; i < db.NoCA; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(db.Timeout))
		db.Conn, err = pgx.Connect(ctx, db.pgUrl)
		cancel()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		db.logger.Error("failed to connect to db", "error", err)
		return err
	}
	db.logger.Info("connected to db")
	return nil
}

func (db *DB) Close() error {
	err := db.Conn.Close(context.Background())
	if err != nil {
		return err
	}
	return nil
}
