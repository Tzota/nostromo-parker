.ONESHELL:
MAKEFLAGS += --silent

clean:
	go clean
	rm ./cmd/terminal ./cover.out
.PHONY: clean

help:
	go help
.PHONY: help

tests:
	go test ./...
.PHONY: tests

coverage:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out
.PHONY: coverage

coverage_raw:
	go test ./... -coverprofile cover.out
	grep -v -e ' 1$$' cover.out
.PHONY: coverage_raw
