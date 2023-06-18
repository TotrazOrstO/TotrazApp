package product

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

	Repository struct {
		db *sqlx.DB
	}
)

func newQueries(db dbtx) *queries {
	return &queries{db: db}
}

func NewProductRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (t *Repository) Close() error {
	return t.db.Close()
}
