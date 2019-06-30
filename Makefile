.PHONY: all build clean deploy test

stage ?= dev

all: deploy test

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/connect connect/connect.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo echo/echo.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy -s $(stage)

test:
	WS_ENDPOINT=$$(sls info -s $(stage) | grep wss | xargs echo) \
	go test -v -count=1 github.com/studzien/sls-ws-example/test
