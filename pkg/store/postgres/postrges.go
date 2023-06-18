package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"totraz_store/pkg/config"
)

func New(cfg config.Postgres) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Driver, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode)
	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres db open fail: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("postgres db ping fail: %w", err)
	}

	return db, nil
}
