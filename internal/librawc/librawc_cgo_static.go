//go:build cgo && libraw_static

package librawc

/*
#cgo linux pkg-config: libraw
#cgo darwin,arm64 CFLAGS: -I/opt/homebrew/opt/libraw/include
#cgo darwin,arm64 LDFLAGS: /opt/homebrew/opt/libraw/lib/libraw.a /opt/homebrew/opt/jpeg-turbo/lib/libjpeg.a /opt/homebrew/opt/little-cms2/lib/liblcms2.a /opt/homebrew/opt/libomp/lib/libomp.a -lz -lc++
#cgo darwin,amd64 CFLAGS: -I/usr/local/opt/libraw/include
#cgo darwin,amd64 LDFLAGS: /usr/local/opt/libraw/lib/libraw.a /usr/local/opt/jpeg-turbo/lib/libjpeg.a /usr/local/opt/little-cms2/lib/liblcms2.a /usr/local/opt/libomp/lib/libomp.a -lz -lc++
*/
import "C"
