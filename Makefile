.PHONY: build example test cover cover-html lint vet fmt clean check

# _example is skipped by ./..., so compile-check it explicitly.
build:
	go build ./...
	go build -o /dev/null ./_example/

# Develops the bundled sample RAW to RAW_CANON_6D.jpg (run from the repo root).
example:
	go run ./_example

test:
	go test -v -cover -coverpkg=./... -race ./...

# Generate a coverage profile and print function-level coverage.
cover:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Generate a coverage profile and open it in a browser.
cover-html:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

vet:
	go vet ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run

clean:
	rm -f coverage.out coverage.html testdata/*.jpg

check: build vet lint test
