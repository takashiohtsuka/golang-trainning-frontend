build:
	go build -o bin/main ./cmd/app/main.go

start:
	air

test:
	go test ./...

.PHONY: build start test
