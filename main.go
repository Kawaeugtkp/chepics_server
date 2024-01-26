package main

import (
	"context"
	"log"

	"github.com/Kawaeugtkp/chepics_server/api"
	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/vanng822/go-solr/solr"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5433/chepics_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannout connect to db:", err)
	}

	store := db.NewStore(conn)
	si, _ := solr.NewSolrInterface("http://localhost:8984/solr", "post")
	server := api.NewServer(store, si)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
