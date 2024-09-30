package store

import (
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Load opens the database at the given path. If `path == ""` then a new in-memory database is returned.
// InitDB initialize the database.
func InitDB(path string) *gorm.DB {
	// Default to an in-memory database.
	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	if path == "" {
		path = "file::memory:?cache=shared"
	}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		slog.Error("cannot open db connection", "err", err)
	}

	// Auto Migrate the schemas
	err = db.AutoMigrate(
		&Player{},
	)
	if err != nil {
		slog.Error("failed to migrate database to current schema: %w", err)
	}

	slog.Info("Database connected", "address", path)

	slog.Info("Database initialized")

	return db
}
