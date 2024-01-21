postgres:
	docker run --name postgresLatestChepics -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgresLatestChepics createdb --username=root --owner=root chepics_db

dropdb:
	docker exec -it postgresLatestChepics dropdb chepics_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/chepics_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/chepics_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server