.PHONY: build example test cover cover-html libraw-check lint vet fmt clean check

# _example is skipped by ./..., so compile-check it explicitly.
build:
	go build ./...
	go build -o /dev/null ./_example/

# Verify cgo can find and link LibRaw, then print the linked runtime version.
libraw-check:
	@echo "cgo: $$(go env CGO_ENABLED)"
	@if [ "$$(go env CGO_ENABLED)" != "1" ]; then \
		echo "error: CGO_ENABLED=1 is required to build go-libraw"; \
		exit 1; \
	fi
	@if command -v pkg-config >/dev/null 2>&1; then \
		if pkg-config --exists libraw; then \
			echo "pkg-config libraw: $$(pkg-config --modversion libraw)"; \
			echo "pkg-config flags: $$(pkg-config --cflags --libs libraw)"; \
		else \
			echo "error: pkg-config is installed, but libraw.pc was not found"; \
			echo "install LibRaw development files or set PKG_CONFIG_PATH to the directory containing libraw.pc"; \
			exit 1; \
		fi; \
	elif [ -d /opt/homebrew/opt/libraw ] || [ -d /usr/local/opt/libraw ]; then \
		echo "pkg-config not found; using Homebrew LibRaw fallback path"; \
	else \
		echo "error: pkg-config was not found and no Homebrew LibRaw fallback path exists"; \
		echo "install pkg-config/pkgconf and LibRaw development files, or see docs/libraw-build.md"; \
		exit 1; \
	fi
	@go test -v -run TestLinkedVersion ./internal/librawc

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

check: libraw-check build vet lint test
