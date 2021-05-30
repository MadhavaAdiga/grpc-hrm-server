package db

import (
	"context"
	"database/sql"
)

// SQLDBTX is an interface common method used in DB and TX struct od database/sql
// implementing will give access to db conn
type SQLDBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// Implements Store interface
type SQLStore struct {
	db SQLDBTX
}

// constructor function
func NewSQlStore(db SQLDBTX) *SQLStore {
	return &SQLStore{
		db: db,
	}
}
