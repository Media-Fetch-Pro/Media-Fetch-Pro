package db

import (
	"context"
	"database/sql"
)

type DB struct {
	// sqlite db connection instance
	DBInstance *sql.DB
}

func NewDB() *DB {
	db := &DB{}
	return db
}

func (db *DB) Open(ctx context.Context) (err error) {
	// db.DBInstance, err = sql.Open("sqlite3", "./foo.db")
	// if err != nil {
	// 	return err
	// }
	return nil
}
