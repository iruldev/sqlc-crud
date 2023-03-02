package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iruldev/sqlc-crud/api"
	database "github.com/iruldev/sqlc-crud/database/sqlc"
	"github.com/iruldev/sqlc-crud/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(fmt.Sprintf("0.0.0.0:%s", config.Port))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
