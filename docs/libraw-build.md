# LibRaw Build Setup

`go-libraw` uses cgo and links against the system LibRaw library. Run this
check after installing LibRaw:

```sh
make libraw-check
```

The check verifies that cgo is enabled, reports the LibRaw discovery path, and
runs the linked runtime version smoke test.

For CI coverage and supported platform assumptions, see
[Support Matrix](support-matrix.md).

## macOS Homebrew

Install LibRaw:

```sh
brew install libraw
```

The current cgo bridge checks the standard Homebrew prefixes:

- Apple Silicon: `/opt/homebrew/opt/libraw`
- Intel: `/usr/local/opt/libraw`

`pkg-config` is useful but not required on macOS while these Homebrew paths are
available.

### Static Linking With `libraw_static`

Default macOS builds link LibRaw dynamically. To opt into the static Homebrew
archives for LibRaw and its open-source dependencies, build with the
`libraw_static` tag:

```sh
go build -tags libraw_static ./...
make libraw-check-static
```

The tag uses the standard Homebrew archive locations for the current
architecture:

```text
/opt/homebrew/opt/libraw/lib/libraw.a
/opt/homebrew/opt/jpeg-turbo/lib/libjpeg.a
/opt/homebrew/opt/little-cms2/lib/liblcms2.a
/opt/homebrew/opt/libomp/lib/libomp.a
```

On Intel macOS the same paths are expected under `/usr/local/opt`. System
libraries such as `libz`, `libc++`, and `libSystem` remain dynamic, which is the
normal macOS linking model. The `libraw-check-static` target builds and runs the
linked-version smoke test, then inspects the test binary with `otool -L` to make
sure no Homebrew dynamic library paths remain.

If Homebrew is installed in a non-standard prefix, use that prefix from the
consumer build system through `CGO_CFLAGS`/`CGO_LDFLAGS` and make sure the
standard archive paths used by the tag are also resolvable. cgo appends
environment flags to package directives; it does not remove directives already
declared by this package.

## Debian And Ubuntu

Install LibRaw and pkg-config metadata:

```sh
sudo apt-get update
sudo apt-get install -y libraw-dev pkg-config
```

Linux builds use:

```sh
pkg-config --cflags --libs libraw
```

If LibRaw is installed in a custom prefix, set `PKG_CONFIG_PATH` to the
directory containing `libraw.pc`.

Static Linux builds are not implemented as a first-class target yet. For local
experiments, use pkg-config's static metadata and any distribution-specific
transitive dependencies:

```sh
PKG_CONFIG="pkg-config --static" go build -tags libraw_static ./...
```

## Fedora

Install LibRaw and pkg-config metadata:

```sh
sudo dnf install LibRaw-devel pkgconf-pkg-config
```

Verify discovery with:

```sh
pkg-config --modversion libraw
make libraw-check
```

## Source Builds

When building LibRaw from source, install it into a prefix that also installs
`libraw.pc`, then point pkg-config at that file:

```sh
export PKG_CONFIG_PATH=/path/to/libraw/lib/pkgconfig:$PKG_CONFIG_PATH
pkg-config --cflags --libs libraw
make libraw-check
```

For unusual source layouts without pkg-config metadata, pass flags through the
standard Go cgo environment variables:

```sh
export CGO_CFLAGS="-I/path/to/libraw/include"
export CGO_LDFLAGS="-L/path/to/libraw/lib -lraw"
go test ./...
```

Prefer pkg-config when possible because it captures platform-specific include
and linker flags in one place.
