MAKEFLAGS += --silent

clean:
	go clean
	rm ./cmd/terminal ./cover.out

help:
	go help

tests:
	go test ./...

coverage:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

coverage_raw:
	go test ./... -coverprofile cover.out
	grep -v -e ' 1$$' cover.out

runc_gc:
	cd cmd/stgc
	make KEY=$(KEY) run

# cover () {
#     t="/tmp/go-cover.$$.tmp"
#     go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
# }

# go test -v -coverprofile cover.out ./YOUR_CODE_FOLDER/...
# go tool cover -html=cover.out -o cover.html
# open cover.html
