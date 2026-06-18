.PHONY: build build-stub clean vet tidy run test

BINARY=nagare

build:
	go build -o bin/$(BINARY) ./cmd/nagare

build-stub:
	go build -tags stub -o bin/$(BINARY) ./cmd/nagare

clean:
	rm -rf bin/

vet:
	go vet ./...

tidy:
	go mod tidy

test:
	go test ./... -count=1

run: build
	./bin/$(BINARY)

run-stub: build-stub
	./bin/$(BINARY)
