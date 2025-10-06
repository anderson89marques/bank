package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTests(t *testing.T) (*sql.DB, error) {
	ctx := context.Background()

	// 1. Start a PostgreSQL container.
	dbContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForListeningPort("5432/tcp")),
	)
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	// 2. Clean up the container after the test finishes.
	t.Cleanup(func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %v", err)
		}
	})

	// 3. Get the connection string for the container.
	connStr, err := dbContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	// 4. Run database migrations
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(filepath.Dir(b))
	migrationsPath := fmt.Sprintf("file://%s", filepath.Join(path, "migrations"))
	m, err := migrate.New(migrationsPath, connStr)
	if err != nil {
		t.Fatalf("failed to create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		t.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully!")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error to connect to database: %w", err)
	}
	return db, err
}
