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

# first git initialization
gitint:
	build: 
		bash: "./bash/commit.sh"

gitcommit:
# Add the file to the staging area
	git add .

# Show the status of the repository
	git status

# Prompt the user for a commit message
	read -p "Enter a commit message: " message

# Create a commit with the given message
	git commit -m "$message"

# Push the commit to the remote branch
	git push origin main


.PHONY: postgres createdb dropdbs migrateup migratedown sqlc gitint gitcommit