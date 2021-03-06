.ONESHELL:
MAKEFLAGS += --silent

clean:
	go clean
	rm ./cover.out
	rm -rf ./cmd/*/out
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
	./arm/release_terminal.sh
.PHONY: release_terminal

deploy_terminal_scp:
	cd build/package/scp
	./terminal.sh pi@192.168.1.104:~/Programs/nostromo-parker-terminal/
.PHONY: deploy_scp

release_redis_stream:
	cd scripts
	./arm/release_redis_stream.sh
.PHONY: release_redis_stream

deploy_redis_stream_scp:
	cd build/package/scp
	./redis_stream.sh pi@192.168.1.104:~/Programs/nostromo-parker-redis-stream/
.PHONY: deploy_redis_stream_scp

build_amd_stream_docker:
	cd scripts/build/
	./docker-clean.sh -a amd64
	./release_redis_stream.sh -a amd64
	./docker-build.sh -a amd64
.PHONY: build_amd_stream_docker

build_arm_stream_docker:
	cd scripts/build/
	./docker-clean.sh -a arm
	GOARM=7 ./release_redis_stream.sh -a arm
	./docker-build.sh -a arm
.PHONY: build_arm_stream_docker
