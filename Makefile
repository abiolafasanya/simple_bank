postgres:
	docker run --name postgres-server -e POSTGRES_USER=abiola -e POSTGRES_PASSWORD=abiola78 -d postgres:15-alpine

# create a new database
createdb: 
	docker exec -it postgres-server createdb --username=abiola --owner=abiola simple_bank

# drop the old database
dropdb: 
	docker exec -it postgres-server dropdb simple_bank

# create or update the migration
migrateup:
	migrate -path db/migration -database "postgresql://abiola:abiola78@localhost:5432/simple_bank?sslmode=disable" -verbose up

#remove or destroy the migration
migratedown:
	migrate -path db/migration -database "postgresql://abiola:abiola78@localhost:5432/simple_bank?sslmode=disable" -verbose down

# sqlc query
sqlc:
	sqlc generate


.PHONY: postgres createdb dropdbs migrateup migratedown sqlc