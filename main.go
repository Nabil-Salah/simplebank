package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer conn.Close()
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
