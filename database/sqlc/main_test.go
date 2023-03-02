package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iruldev/sqlc-crud/util"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = "root:root@tcp(127.0.0.1:3306)/learn?parseTime=true"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
