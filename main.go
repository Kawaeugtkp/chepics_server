package main

import (
	"database/sql"
	"log"

	"github.com/Kawaeugtkp/chepics_server/api"
	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	_ "github.com/lib/pq"
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/chepics_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannout connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}