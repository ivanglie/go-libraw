.PHONY: api-inventory api-coverage check-api-inventory check-api-coverage generate build build-static examples test test-fast test-fixtures cover cover-html libraw-check libraw-check-static lint vet fmt clean check release-check

LIBRAW_HEADERS ?= testdata/headers/libraw

build:
	go build ./...

build-static:
	go build -tags libraw_static ./...

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

# Verify the opt-in macOS static link path and assert Homebrew dylibs are absent.
libraw-check-static:
	@echo "cgo: $$(go env CGO_ENABLED)"
	@if [ "$$(go env CGO_ENABLED)" != "1" ]; then \
		echo "error: CGO_ENABLED=1 is required to build go-libraw"; \
		exit 1; \
	fi
	@if [ "$$(go env GOOS)" != "darwin" ]; then \
		echo "error: libraw-check-static currently supports macOS only"; \
		echo "Linux static linking can be tested with PKG_CONFIG=\"pkg-config --static\" and project-specific CGO_LDFLAGS"; \
		exit 1; \
	fi
	@if [ "$$(go env GOARCH)" = "arm64" ]; then \
		prefix="/opt/homebrew/opt"; \
	else \
		prefix="/usr/local/opt"; \
	fi; \
	for archive in \
		"$$prefix/libraw/lib/libraw.a" \
		"$$prefix/jpeg-turbo/lib/libjpeg.a" \
		"$$prefix/little-cms2/lib/liblcms2.a" \
		"$$prefix/libomp/lib/libomp.a"; do \
		if [ ! -f "$$archive" ]; then \
			echo "error: required static archive not found: $$archive"; \
			exit 1; \
		fi; \
	done
	go test -v -tags libraw_static -run TestLinkedVersion ./internal/librawc
	@mkdir -p tmp
	go test -c -tags libraw_static -o tmp/librawc-static.test ./internal/librawc
	@if otool -L tmp/librawc-static.test | grep -E '/opt/homebrew|/usr/local'; then \
		echo "error: static test binary still links Homebrew dynamic libraries"; \
		exit 1; \
	fi

# Regenerate the LibRaw API inventory and coverage map.
api-inventory:
	go run ./cmd/libraw-api-inventory -headers "$(LIBRAW_HEADERS)" -update-coverage

api-coverage:
	go run ./cmd/libraw-api-inventory -headers "$(LIBRAW_HEADERS)" -coverage-report docs/libraw-api-coverage.md

# Regenerate Go files from checked-in LibRaw fixture headers.
generate:
	go run ./tools/gen-constants
	gofmt -w constants_generated.go

# Verify committed LibRaw API inventory and coverage map are current.
check-api-inventory:
	go run ./cmd/libraw-api-inventory -headers "$(LIBRAW_HEADERS)" -check

check-api-coverage:
	go run ./cmd/libraw-api-inventory -headers "$(LIBRAW_HEADERS)" -check -coverage-report docs/libraw-api-coverage.md

# Run the LibRaw sample-parity examples on bundled fixtures.
# Outputs land under tmp/outputs and tmp/examples and are removed by `make clean`.
# See docs/examples.md for the upstream sample each one mirrors.
examples:
	@mkdir -p tmp/outputs tmp/examples
	go run ./cmd/simple-dcraw testdata/RAW_CANON_R6.CR3
	go run ./cmd/raw-identify testdata/RAW_CANON_R6.CR3
	go run ./cmd/raw-textdump testdata/RAW_NIKON_ZFC.NEF
	go run ./cmd/mem-image testdata/RAW_RICOH_GR3X.DNG
	go run ./cmd/thumb-extract testdata/RAW_CANON_R6.CR3
	go run ./cmd/unprocessed-raw testdata/RAW_RICOH_GR3X.DNG
	go run ./cmd/openbayer
	go run ./cmd/multirender testdata/RAW_RICOH_GR3X.DNG
	go run ./cmd/four-channels testdata/RAW_RICOH_GR3X.DNG

test:
	go test -v -cover -coverpkg=./... -race ./...

test-fast:
	go test -short ./...

test-fixtures:
	go test -count=1 -run 'Test(FixtureRegression|MetadataForFixtures|MakerNotesForFixtures)' ./...

# Generate a coverage profile and print function-level coverage.
# cmd/ and tools/ are excluded from the denominator: they are main packages
# with no tests and would artificially deflate the library coverage number.
cover:
	go test -coverpkg=./internal/...,./pkg/... -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Generate a coverage profile and open it in a browser.
cover-html:
	go test -coverpkg=./internal/...,./pkg/... -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

vet:
	go vet ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run

clean:
	rm -f coverage.out coverage.html testdata/*.jpg
	rm -rf tmp/outputs tmp/examples

check: libraw-check check-api-inventory build vet lint test

release-check: check check-api-coverage test-fixtures examples
	$(MAKE) clean
