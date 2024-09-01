package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/popeskul/awesome-messanger/services/platform/database/postgres"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/config"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/querybuilder"
)

func setupTestContainer(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithStartupTimeout(2 * time.Minute),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
		},
	}

	t.Log("Starting PostgreSQL container...")
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	t.Log("Container started, getting host...")
	ip, err := pgContainer.Host(ctx)
	require.NoError(t, err)
	t.Logf("Container IP: %s", ip)

	t.Log("Getting mapped port...")
	port, err := pgContainer.MappedPort(ctx, "5432")
	require.NoError(t, err)
	t.Logf("Mapped port: %s", port.Port())

	connectionString := fmt.Sprintf("postgres://user:password@%s:%s/testdb", ip, port.Port())
	t.Logf("Connection string: %s", connectionString)

	t.Log("Waiting for database to be ready...")
	err = waitForDatabase(ctx, connectionString)
	require.NoError(t, err)

	t.Log("Database is ready")

	return pgContainer, connectionString
}

func waitForDatabase(ctx context.Context, connectionString string) error {
	var err error
	for i := 0; i < 60; i++ {
		config, err := pgxpool.ParseConfig(connectionString)
		if err != nil {
			return err
		}
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err == nil {
			defer pool.Close()
			err = pool.Ping(ctx)
			if err == nil {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("database not ready after 60 seconds: %v", err)
}

func setupTestDB(t *testing.T) (ports.Connection, func()) {
	t.Log("Setting up test container...")
	container, connectionString := setupTestContainer(t)

	t.Log("Creating database configuration...")
	cfg := config.New().
		WithConnectionString(connectionString).
		WithMaxConnections(5)

	ctx := context.Background()
	t.Log("Creating new database connection...")
	db, err := postgres.NewDatabase(ctx, cfg)
	require.NoError(t, err)
	t.Log("Database connection created successfully")

	return db, func() {
		t.Log("Closing database connection...")
		db.Close()
		t.Log("Terminating container...")
		container.Terminate(ctx)
	}
}

func TestNewDatabase(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	err := db.Ping(ctx)
	assert.NoError(t, err)
}

func TestConnection_Exec(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	_, err := db.Exec(ctx, "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name TEXT)")
	assert.NoError(t, err)

	_, err = db.Exec(ctx, "INSERT INTO test_table (name) VALUES ($1)", "test_name")
	assert.NoError(t, err)

	// Clean up
	_, err = db.Exec(ctx, "DROP TABLE test_table")
	assert.NoError(t, err)
}

func TestConnection_Query(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	_, err := db.Exec(ctx, "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name TEXT)")
	require.NoError(t, err)

	_, err = db.Exec(ctx, "INSERT INTO test_table (name) VALUES ($1)", "test_name")
	require.NoError(t, err)

	rows, err := db.Query(ctx, "SELECT name FROM test_table")
	assert.NoError(t, err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		assert.NoError(t, err)
		assert.Equal(t, "test_name", name)
	}

	// Clean up
	_, err = db.Exec(ctx, "DROP TABLE test_table")
	assert.NoError(t, err)
}

func TestConnection_QueryRow(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	_, err := db.Exec(ctx, "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name TEXT)")
	require.NoError(t, err)

	_, err = db.Exec(ctx, "INSERT INTO test_table (name) VALUES ($1)", "test_name")
	require.NoError(t, err)

	var name string
	err = db.QueryRow(ctx, "SELECT name FROM test_table WHERE id = $1", 1).Scan(&name)
	assert.NoError(t, err)
	assert.Equal(t, "test_name", name)

	// Clean up
	_, err = db.Exec(ctx, "DROP TABLE test_table")
	assert.NoError(t, err)
}

func TestTransaction(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	_, err := db.Exec(ctx, "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name TEXT)")
	require.NoError(t, err)

	tx, err := db.BeginTx(ctx)
	require.NoError(t, err)

	_, err = tx.Exec(ctx, "INSERT INTO test_table (name) VALUES ($1)", "test_name")
	assert.NoError(t, err)

	err = tx.Commit(ctx)
	assert.NoError(t, err)

	var name string
	err = db.QueryRow(ctx, "SELECT name FROM test_table WHERE id = $1", 1).Scan(&name)
	assert.NoError(t, err)
	assert.Equal(t, "test_name", name)

	// Test rollback
	tx, err = db.BeginTx(ctx)
	require.NoError(t, err)

	_, err = tx.Exec(ctx, "INSERT INTO test_table (name) VALUES ($1)", "rollback_name")
	assert.NoError(t, err)

	err = tx.Rollback(ctx)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM test_table").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count) // Should still be 1 after rollback

	// Clean up
	_, err = db.Exec(ctx, "DROP TABLE test_table")
	assert.NoError(t, err)
}

func TestQueryBuilder(t *testing.T) {
	builder := querybuilder.NewSelectBuilder("users", "id", "name").
		Where("age > $1", 18)

	query, args, err := builder.Build()
	assert.NoError(t, err)
	assert.Equal(t, "SELECT id, name FROM users WHERE age > $1", query)
	assert.Equal(t, []interface{}{18}, args)
}
