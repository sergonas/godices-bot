.PHONY: migrate
migrate:
	migrate -path ./db/migrations/ -database postgres://postgres:example@localhost:5432/postgres?sslmode=disable up
.PHONY: build
build: 
	go build -v ./cmd/godicesbot

.DEFAULT_GOAL := build