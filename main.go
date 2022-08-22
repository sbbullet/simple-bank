package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sbbullet/simple-bank/api"
	db "github.com/sbbullet/simple-bank/db/sqlc"
	"github.com/sbbullet/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database. Error: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server. Error: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create server. Error: ", err)
	}
}
