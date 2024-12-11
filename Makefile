DATABASE_URL = "postgres://localhost:5432/syncly?user=postgres&password=postgres"


build:
	docker build -t syncly -f ./docker/Dockerfile .

run:
	docker-compose up

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down
