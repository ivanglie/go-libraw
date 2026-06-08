package libraw

import (
	"errors"
	"strings"
	"testing"
)

func TestErrorCodeString(t *testing.T) {
	msg := StrError(LIBRAW_FILE_UNSUPPORTED)
	if strings.TrimSpace(msg) == "" {
		t.Fatal("StrError returned empty message")
	}
}

func TestToError(t *testing.T) {
	if err := ToError(LIBRAW_SUCCESS); err != nil {
		t.Fatalf("ToError(LIBRAW_SUCCESS) = %v, want nil", err)
	}

	err := ToError(LIBRAW_FILE_UNSUPPORTED)
	if err == nil {
		t.Fatal("ToError(LIBRAW_FILE_UNSUPPORTED) = nil")
	}

	var librawErr Error
	if !errors.As(err, &librawErr) {
		t.Fatalf("error %T does not unwrap as Error", err)
	}
	if librawErr.Code != LIBRAW_FILE_UNSUPPORTED {
		t.Fatalf("Code = %d, want %d", librawErr.Code, LIBRAW_FILE_UNSUPPORTED)
	}
	if !strings.Contains(err.Error(), "-2") {
		t.Fatalf("Error() = %q, want numeric code", err.Error())
	}
}

func TestProgressString(t *testing.T) {
	msg := LIBRAW_PROGRESS_OPEN.String()
	if strings.TrimSpace(msg) == "" {
		t.Fatal("Progress.String returned empty message")
	}
}

func TestStrProgress(t *testing.T) {
	msg := StrProgress(LIBRAW_PROGRESS_START)
	if strings.TrimSpace(msg) == "" {
		t.Fatal("StrProgress returned empty message")
	}
}

func TestGeneratedVersionConstants(t *testing.T) {
	if LIBRAW_MAJOR_VERSION != 0 {
		t.Fatalf("LIBRAW_MAJOR_VERSION = %d, want 0", LIBRAW_MAJOR_VERSION)
	}
	if LIBRAW_MINOR_VERSION <= 0 {
		t.Fatalf("LIBRAW_MINOR_VERSION = %d, want positive", LIBRAW_MINOR_VERSION)
	}
	if LIBRAW_VERSION <= 0 {
		t.Fatalf("LIBRAW_VERSION = %d, want positive", LIBRAW_VERSION)
	}
}
