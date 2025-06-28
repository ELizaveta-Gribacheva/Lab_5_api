run:
	go run cmd/api/main.go

test:
	go test ./... -v

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose down 1

.PHONY: run test migrateup migrateup1 migratedown migratedown1
