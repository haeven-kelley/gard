// internal/storage/storage.go
package storage

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Initialize initializes the database connection
func Initialize() *bun.DB {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

// SaveMPDInfo saves information about the MPD file to the database
func SaveMPDInfo(filePath string) error {
	// Implement your database logic here
	return nil
}