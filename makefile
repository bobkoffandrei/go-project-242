build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

.PHONY: test
test:
	go test -v ./testdata

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix