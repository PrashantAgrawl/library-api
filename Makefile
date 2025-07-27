run:
	docker-compose up --build

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/library?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/library?sslmode=disable" down
