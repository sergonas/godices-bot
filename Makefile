.PHONY: build
build: go build -v ./cmd/godicesbot

.DEFAULT_GOAL := build