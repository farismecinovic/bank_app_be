package main

import (
	"database/sql"
	"log"

	"github.com/farismecinovic/bankapp/api"
	db "github.com/farismecinovic/bankapp/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:swordfish@localhost:5432/simple_bank?sslmode=disable"
	serverAdress = "0.0.0.0:5050"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to the DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("Can't start the server!", err)
	}
}
