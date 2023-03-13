migrateup:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose up

migrateup1:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose up 1

migratedown:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose down

migratedown1:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination database/mock/store.go github.com/iruldev/sqlc-crud/database/sqlc Store

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc test server mock