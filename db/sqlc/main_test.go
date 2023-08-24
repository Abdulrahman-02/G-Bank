package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// For now I will declare them as consts but later on I will change them to env variables
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
