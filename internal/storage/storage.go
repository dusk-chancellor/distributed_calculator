package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Model functions for initializing database and its tables

type Storage struct {
	db *sql.DB
}

// Methods for interactions with database
type ExpressionInteractor interface { 
	InsertExpression(ctx context.Context, expr *Expression) (int64, error)
	SelectExpressionsByID(ctx context.Context, userID int64) ([]Expression, error)
	DeleteExpression(ctx context.Context, id int64) error
	UpdateExpression(ctx context.Context, answer, status string, id int64,) error
	SelectAllExpressions(ctx context.Context) ([]Expression, error)
}

// New - creates new database in the given storagePath with tables
func New(storagePath string) (*Storage, error) {
	const op = "storage/storage-New"
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	if err = createTables(ctx, db); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &Storage{db: db}, nil
}

// createTables - pops up two tables in the given database
func createTables(ctx context.Context, db *sql.DB) error {
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

				FOREIGN KEY (userid) REFERENCES users (id)
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