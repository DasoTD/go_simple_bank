package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
	_"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

const (
	DBDriver ="postgres"
	DBSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	testDB, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}