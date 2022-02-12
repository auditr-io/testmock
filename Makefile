.PHONY: build test clean

build:
	export GO111MODULE=on
	go build -ldflags="-s -w" -o bin/testmock *.go

test:
	go test -timeout 30s -v ./...

clean:
	rm -rf ./bin ./vendor Gopkg.lock
