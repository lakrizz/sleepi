package db

import (
	"database/sql"
	"embed"
	"log/slog"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite" // or mattn/go-sqlite3
)

//go:embed migrations/*.sql
var fs embed.FS

func RunMigrations(db *sql.DB) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	slog.Info("migrations applied")
	return nil
}
