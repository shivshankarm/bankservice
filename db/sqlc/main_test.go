package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/shivshankarm/bankservice/util"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(" Cannot connect to DB: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
