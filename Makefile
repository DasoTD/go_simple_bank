postgres:
	sudo docker run -d --name postgres15alpl -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine3.17
creatdb:
	sudo docker exec -it postgres15alpl createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres15alpl dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc


server:
	go run main.go


mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/dasotd/go_simple_bank/db/sqlc Store
.PHONY: creatdb dropdb postgres migratedown migrateup test server mock