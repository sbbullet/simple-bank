package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sbbullet/simple-bank/api"
	db "github.com/sbbullet/simple-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database. Error: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot create server. Error: ", err)
	}
}
