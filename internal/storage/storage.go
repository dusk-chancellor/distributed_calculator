package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	if err = CreateTables(ctx, db); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &Storage{db: db}, nil
}

func CreateTables(ctx context.Context, db *sql.DB) error {
	const (
		usersTable = `
			CREATE TABLE IF NOT EXISTS users(
				id 		 INTEGER PRIMARY KEY AUTOINCREMENT,
				name	 TEXT NOT NULL UNIQUE,
				password TEXT NOT NULL
			);`

		expressionsTable = `
			CREATE TABLE IF NOT EXISTS expressions(
				id 		   INTEGER PRIMARY KEY AUTOINCREMENT,
				userid 	   INTEGER,
				expression TEXT NOT NULL,
				answer 	   TEXT,
				date 	   TEXT,
				status 	   TEXT,

				FOREIGN KEY (userid) REFERENCES users(id)
			);`
	)

	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, expressionsTable); err != nil {
		return err
	}

	return nil
}