ROOT=cmd/currency-task/main.go
MIG=cmd/migrator/main.go


run:
	go run ${ROOT}

mig:
	go run ${MIG}

start:
	docker compose up -d --build

down:
	docker compose down