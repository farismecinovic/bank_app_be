package main

import (
	"database/sql"
	"log"

	"github.com/farismecinovic/bankapp/api"
	db "github.com/farismecinovic/bankapp/db/sqlc"
	"github.com/farismecinovic/bankapp/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config!", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to the DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start the server!", err)
	}
}
