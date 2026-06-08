# LibRaw Build Setup

`go-libraw` uses cgo and links against the system LibRaw library. Run this
check after installing LibRaw:

```sh
make libraw-check
```

The check verifies that cgo is enabled, reports the LibRaw discovery path, and
runs the linked runtime version smoke test.

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

