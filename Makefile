GO = go

export_envs:
	@export $(cat .env | xargs)

db_login:
	psql ${DATABASE_URL}	

# RUN LOCAL
.PHONY: build-api  run-dev

build-api:
	$(GO) build -o /cmd/api/api ./cmd/api/main.go

run-dev:
	$(GO) run cmd/api/main.go

test:
	$(GO) test ./...

.PHONY: run-db run-api
# Run project
run:
	docker-compose up --build db -d && \
	docker-compose up --build --force-recreate api -d

# Run api
run-api:
	docker-compose up --build --force-recreate api -d

# Run db
run-db:
	docker-compose up --build db -d

# Database migrations
migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrate:
	migrate -database ${DATABASE_URL} -path migrations up

# Swagger
swag:
	swag init ./cmd/api/main.go -o docs --parseDependency --parseInternal
