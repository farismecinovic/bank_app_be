package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/farismecinovic/bankapp/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, configError := util.LoadConfig("../../")
	if configError != nil {
		log.Fatal("Can't load the config!: ", configError)
	}

	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to the DB: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
