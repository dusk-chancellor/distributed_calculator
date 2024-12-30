package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dusk-chancellor/distributed_calculator/sso/internal/config"
	"github.com/dusk-chancellor/distributed_calculator/sso/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func New(ctx context.Context, cfg config.Database) (*sql.DB, error) {
	dsn := buildDSN(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetBaseFS(migrations.Migrations.FS)

	if err := goose.Up(db, migrations.Migrations.Dir); err != nil {
		return err
	}

	return nil
}

func buildDSN(cfg config.Database) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)
}
