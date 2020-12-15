package dbutil

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pachyderm/pachyderm/src/client/pkg/require"
)

// set this to true if you want to keep the database around
var devDontDropDatabase = false

const (
	// DefaultPostgresHost for tests
	DefaultPostgresHost = "127.0.0.1"
	// DefaultPostgresPort for tests
	DefaultPostgresPort = 32228
	// TestPostgresUser is the default postgres user
	TestPostgresUser = "postgres"
)

// NewTestDB connects to postgres using the default settings, creates a database with a unique name
// then calls cb with a sqlx.DB configured to use the newly created database.
// After cb returns the database is dropped.
func NewTestDB(t testing.TB) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", DefaultPostgresHost, DefaultPostgresPort, TestPostgresUser)
	db := sqlx.MustOpen("postgres", dsn)
	dbName := ephemeralDBName()
	t.Log("creating database:", dbName)
	db.MustExec(`CREATE DATABASE ` + dbName)
	t.Cleanup(func() {
		if !devDontDropDatabase {
			db.MustExec("DROP DATABASE " + dbName)
		}
		require.NoError(t, db.Close())
	})
	db2 := sqlx.MustConnect("postgres", dsn+" dbname="+dbName)
	t.Log("database", dbName, "successfully created")
	t.Cleanup(func() {
		require.NoError(t, db2.Close())
	})
	return db2
}

func ephemeralDBName() string {
	// TODO: it looks like postgres is truncating identifiers to 32 bytes,
	// it should be 64 but we might be passing the name as non-ascii, i'm not really sure.
	// for now just use a random int, but it would be nice to go back to names with a timestamp.
	return fmt.Sprintf("test_%08x", rand.Uint64())
	//now := time.Now()
	// test_<date>T<time>_<random int>
	// return fmt.Sprintf("test_%04d%02d%02dT%02d%02d%02d_%04x",
	// 	now.Year(), now.Month(), now.Day(),
	// 	now.Hour(), now.Minute(), now.Second(),
	// 	rand.Uint32())
}

// DBParams are parameters passed to the db constructor.
type DBParams struct {
	Host       string
	Port       int
	User, Pass string
	DBName     string
}

// NewDB returns a db created with the given parameters
func NewDB(x DBParams) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", x.Host, x.Port, x.User, x.Pass, x.DBName)
	return sqlx.Open("postgres", dsn)
}
