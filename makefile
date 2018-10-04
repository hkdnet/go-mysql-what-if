default: all

all: bin/duplicate-entry

test: all
	bin/duplicate-entry

bin/duplicate-entry: cmd/duplicate-entry/main.go
	cd cmd/duplicate-entry && go get -d -v ./...
	go build -o $@ cmd/duplicate-entry/main.go

.PHONY: all test
