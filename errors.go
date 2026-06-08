package libraw

import (
	"fmt"

	"github.com/ivanglie/go-libraw/internal/librawc"
)

// ErrorCode is a LibRaw error code.
type ErrorCode int

// Error describes a LibRaw failure while preserving the original numeric code.
type Error struct {
	Code ErrorCode
}

// Error returns the LibRaw error message with the numeric code.
func (e Error) Error() string {
	return fmt.Sprintf("libraw: %s (%d)", e.Code.String(), int(e.Code))
}

// String returns LibRaw's message for the error code.
func (c ErrorCode) String() string {
	if msg := librawc.StrError(int(c)); msg != "" {
		return msg
	}
	return fmt.Sprintf("unknown LibRaw error %d", int(c))
}

// ToError converts a LibRaw status code into a Go error.
func ToError(code ErrorCode) error {
	if code == 0 {
		return nil
	}
	return Error{Code: code}
}

// StrError returns LibRaw's message for an error code.
func StrError(code ErrorCode) string {
	return code.String()
}

// Progress is a LibRaw progress stage.
type Progress int

// String returns LibRaw's progress name for the stage.
func (p Progress) String() string {
	if msg := librawc.StrProgress(int(p)); msg != "" {
		return msg
	}
	return fmt.Sprintf("unknown LibRaw progress %d", int(p))
}

// StrProgress returns LibRaw's message for a progress stage.
func StrProgress(progress Progress) string {
	return progress.String()
}
