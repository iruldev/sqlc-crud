migrateup:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose up
migratedown:
	migrate -path database/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/learn" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...