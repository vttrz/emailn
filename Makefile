test:
	go test ./...

up:
	docker compose up -d

down:
	docker compose down