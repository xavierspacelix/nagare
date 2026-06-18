.PHONY: build clean vet tidy run

BINARY=nagare

build:
	go build -o bin/$(BINARY) ./cmd/nagare

clean:
	rm -rf bin/

vet:
	go vet ./...

tidy:
	go mod tidy

run: build
	./bin/$(BINARY)
