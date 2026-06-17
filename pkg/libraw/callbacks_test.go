package libraw

import (
	"errors"
	"testing"
)

func TestProgressCallbackReceivesEvents(t *testing.T) {
	p := openProcessor(t)

	var stages []Progress
	if err := p.SetProgressHandler(func(stage Progress, iteration, expected int) int {
		stages = append(stages, stage)
		return 0
	}); err != nil {
		t.Fatalf("SetProgressHandler() error = %v", err)
	}

	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		t.Fatalf("DcrawProcess() error = %v", err)
	}

	if len(stages) == 0 {
		t.Fatal("progress callback received no events")
	}
}

func TestProgressCallbackCancels(t *testing.T) {
	p := openProcessor(t)

	if err := p.SetProgressHandler(func(stage Progress, iteration, expected int) int {
		return 1 // cancel immediately
	}); err != nil {
		t.Fatalf("SetProgressHandler() error = %v", err)
	}

	// Progress events start during identification, so OpenFile is cancelled.
	err := p.OpenFile(sampleRAW)
	if err == nil {
		t.Fatal("OpenFile() with cancelling callback returned nil error")
	}
	var le Error
	if !errors.As(err, &le) || le.Code != LIBRAW_CANCELLED_BY_CALLBACK {
		t.Fatalf("OpenFile() error = %v, want LIBRAW_CANCELLED_BY_CALLBACK", err)
	}
}

func TestProgressCallbackPanicCancels(t *testing.T) {
	p := openProcessor(t)

	if err := p.SetProgressHandler(func(stage Progress, iteration, expected int) int {
		panic("boom")
	}); err != nil {
		t.Fatalf("SetProgressHandler() error = %v", err)
	}

	// A panicking progress callback is recovered and treated as cancellation,
	// so the first triggering call (OpenFile) returns the cancellation error.
	err := p.OpenFile(sampleRAW)
	if err == nil {
		t.Fatal("OpenFile() with panicking callback returned nil error")
	}
	var le Error
	if !errors.As(err, &le) || le.Code != LIBRAW_CANCELLED_BY_CALLBACK {
		t.Fatalf("OpenFile() error = %v, want LIBRAW_CANCELLED_BY_CALLBACK", err)
	}
}

func TestTagParserCallbacks(t *testing.T) {
	p := openProcessor(t)

	var exifTags, makerTags int
	if err := p.SetExifParserHandler(func(tag, typ, length int, order uint32, base int64) {
		exifTags++
	}); err != nil {
		t.Fatalf("SetExifParserHandler() error = %v", err)
	}
	if err := p.SetMakerNotesHandler(func(tag, typ, length int, order uint32, base int64) {
		makerTags++
	}); err != nil {
		t.Fatalf("SetMakerNotesHandler() error = %v", err)
	}

	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	// EXIF parsing happens during identification; at least one handler should
	// observe tags for a mainstream camera file.
	if exifTags == 0 && makerTags == 0 {
		t.Fatal("neither EXIF nor maker-note callback received any tags")
	}
}

func TestDataErrorHandlerDoesNotBreakProcessing(t *testing.T) {
	p := openProcessor(t)

	if err := p.SetDataErrorHandler(func(file string, offset int64) {}); err != nil {
		t.Fatalf("SetDataErrorHandler() error = %v", err)
	}
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
}

func TestNilHandlerIsSafe(t *testing.T) {
	p := openProcessor(t)

	// Register a counting progress handler plus the others, then clear each by
	// passing nil. A nil handler must actually unregister on the LibRaw side, so
	// the counter must stay at zero once processing runs.
	var progressCalls int
	if err := p.SetProgressHandler(func(Progress, int, int) int { progressCalls++; return 0 }); err != nil {
		t.Fatalf("SetProgressHandler() error = %v", err)
	}
	if err := p.SetDataErrorHandler(func(string, int64) {}); err != nil {
		t.Fatalf("SetDataErrorHandler() error = %v", err)
	}
	if err := p.SetExifParserHandler(func(int, int, int, uint32, int64) {}); err != nil {
		t.Fatalf("SetExifParserHandler() error = %v", err)
	}
	if err := p.SetMakerNotesHandler(func(int, int, int, uint32, int64) {}); err != nil {
		t.Fatalf("SetMakerNotesHandler() error = %v", err)
	}

	clear := []struct {
		name string
		fn   func() error
	}{
		{"SetProgressHandler", func() error { return p.SetProgressHandler(nil) }},
		{"SetDataErrorHandler", func() error { return p.SetDataErrorHandler(nil) }},
		{"SetExifParserHandler", func() error { return p.SetExifParserHandler(nil) }},
		{"SetMakerNotesHandler", func() error { return p.SetMakerNotesHandler(nil) }},
	}
	for _, c := range clear {
		if err := c.fn(); err != nil {
			t.Fatalf("%s(nil) error = %v", c.name, err)
		}
	}

	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		t.Fatalf("DcrawProcess() error = %v", err)
	}

	if progressCalls != 0 {
		t.Fatalf("progress handler fired %d times after being cleared with nil", progressCalls)
	}
}

func TestCallbacksSurviveRecycle(t *testing.T) {
	p := openProcessor(t)

	var progress, exif, makernotes int
	if err := p.SetProgressHandler(func(Progress, int, int) int { progress++; return 0 }); err != nil {
		t.Fatalf("SetProgressHandler() error = %v", err)
	}
	if err := p.SetDataErrorHandler(func(string, int64) {}); err != nil {
		t.Fatalf("SetDataErrorHandler() error = %v", err)
	}
	if err := p.SetExifParserHandler(func(int, int, int, uint32, int64) { exif++ }); err != nil {
		t.Fatalf("SetExifParserHandler() error = %v", err)
	}
	if err := p.SetMakerNotesHandler(func(int, int, int, uint32, int64) { makernotes++ }); err != nil {
		t.Fatalf("SetMakerNotesHandler() error = %v", err)
	}

	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("first OpenFile() error = %v", err)
	}
	if err := p.Recycle(); err != nil {
		t.Fatalf("Recycle() error = %v", err)
	}

	beforeProgress := progress
	beforeExif := exif
	// Callbacks must still fire after Recycle, which zeroes LibRaw's C-side
	// handler pointers. reregisterCallbacks must reinstall them.
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("second OpenFile() error = %v", err)
	}
	if progress == beforeProgress {
		t.Fatal("progress callback stopped firing after Recycle")
	}
	// exif OR makernotes must fire (camera-dependent); at least one must increment.
	if exif == beforeExif && makernotes == 0 {
		t.Fatal("neither exif nor makernotes callback fired after Recycle")
	}
}

func TestCallbackHandlersAfterCloseReturnErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"SetProgressHandler":   func() error { return p.SetProgressHandler(nil) },
		"SetDataErrorHandler":  func() error { return p.SetDataErrorHandler(nil) },
		"SetExifParserHandler": func() error { return p.SetExifParserHandler(nil) },
		"SetMakerNotesHandler": func() error { return p.SetMakerNotesHandler(nil) },
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}
