.ONESHELL:
MAKEFLAGS += --silent

clean:
	go clean
	rm ./cover.out

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

release_terminal:
	cd scripts
	./release_terminal.sh
.PHONY: release_terminal

deploy_terminal_scp:
	cd build/package/scp
	./index.sh pi@192.168.1.104:~/Programs/nostromo-parker/
.PHONY: deploy_scp
