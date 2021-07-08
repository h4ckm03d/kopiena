package core

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// DBHandler is an interface that models the standard transaction in
// `database/sql`.
//
// To ensure `TxFn` funcs cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
type DBHandler interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// A Txfn is a function that will be called with an initialized `DBHandler` object
// that can be used for executing statements and queries against a database.
type TxFn func(DBHandler) error

type TxStarter interface {
	Begin() (*sql.Tx, error)
}

type TxHandler interface {
	WithTransaction(TxFn) error
}

func NewTxSQLHandler(db *sql.DB) *TxSQLHandler {
	return &TxSQLHandler{db}
}

type TxSQLHandler struct {
	db *sql.DB
}

func (th *TxSQLHandler) WithTransaction(fn TxFn) (err error) {
	tx, err := th.db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "fail rollback panic: %+v", rollbackErr)
			}
			panic(p)
		}

		if err != nil {
			// something went wrong, rollback
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "fail rollback error: %+v", rollbackErr)
			}
			return
		}

		// all good, commit
		err = errors.Wrap(tx.Commit(), "fail commit")
	}()

	err = fn(tx)
	return
}
