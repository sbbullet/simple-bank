postgres:
	docker run -p 5432:5432 --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migration -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migration -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migration -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down -all

migratedown1:
	docker run --rm -v $(CURDIR)/db/migration:/db/migration --network=host migrate/migrate -path=/db/migration -database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover -count=1 ./...

server:
	go run main.go

mock:
	mockgen --package=mockdb --destination=db/mock/store.go github.com/sbbullet/simple-bank/db/sqlc Store

proto:
	rm -rf pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb \
	--grpc-gateway_opt=paths=source_relative \
  proto/*.proto

evans:
	evans --host localhost --port 8081 -r repl

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 server mock proto evans