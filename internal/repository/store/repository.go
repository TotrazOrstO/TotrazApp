package store

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type (
	dbtx interface {
		ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
		PrepareContext(context.Context, string) (*sql.Stmt, error)
		QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	}

	queries struct {
		db dbtx
	}

	StoreRepository struct {
		db *sqlx.DB
	}
)

func newQueries(db dbtx) *queries {
	return &queries{db: db}
}

func NewStoreRepository(db *sqlx.DB) *StoreRepository {
	return &StoreRepository{
		db: db,
	}
}

func (t *StoreRepository) Close() error {
	return t.db.Close()
}
