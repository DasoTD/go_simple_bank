package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/dasotd/go_simple_bank/util"
	_ "github.com/lib/pq"
	_ "github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB



func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}