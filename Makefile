postgres:
	docker run -p 5432:5432 --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migration -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migrations -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover -count=1 ./...

.PHONY: postgres createdb dropdb migrateup migratedown