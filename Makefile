build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

migrate:
	go run ./cmd/pokes_migrate/main.go

.PHONY: test
test:
	docker compose exec -w /usr/src/app/ app go test -v ./test/

run:
	go run ./cmd/pokes_migrate/main.go
	go run ./cmd/pokewordle_solver
