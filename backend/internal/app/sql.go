package app

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/lakrizz/sleepi/internal/infra/db"
	"github.com/lakrizz/sleepi/internal/repositories"

	_ "github.com/mattn/go-sqlite3"
)

func (a *App) initDatabase() error {
	// open db first
	sqldb, err := sql.Open("sqlite3", "sleepi.db")
	if err != nil {
		return fmt.Errorf("error opening sqlite db: %w", err)
	}
	slog.Info("database opened")

	// run migrations
	err = db.RunMigrations(sqldb)
	if err != nil {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	a.DB = sqldb

	// sqlc queries
	queries := db.New(sqldb)
	a.Queries = queries

	// now repositories
	alarmRepository, err := repositories.NewAlarmsRepository(a.Queries)
	if err != nil {
		return fmt.Errorf("could not initalize alarm repository: %w", err)
	}

	a.AlarmRepository = alarmRepository
	return nil
}
