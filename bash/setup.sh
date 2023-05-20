!#/bin/bash

# Create db migrations
migrate create -ext sql -dir db/migration -seq init_schema

# $ migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2