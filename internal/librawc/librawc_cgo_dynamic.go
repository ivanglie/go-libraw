//go:build cgo && !libraw_static

package librawc

/*
#cgo linux pkg-config: libraw
#cgo darwin,arm64 CFLAGS: -I/opt/homebrew/opt/libraw/include
#cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/libraw/lib -lraw
#cgo darwin,amd64 CFLAGS: -I/usr/local/opt/libraw/include
#cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/libraw/lib -lraw
*/
import "C"
