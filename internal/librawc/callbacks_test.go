//go:build cgo

package librawc

import "testing"

func TestCallbackRegistryReleasedOnClose(t *testing.T) {
	h, err := New(0)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	h.SetProgressCallback(func(stage, iteration, expected int) int { return 0 })
	if h.cb == nil || h.cbHandle == 0 {
		t.Fatal("callback registry was not created on registration")
	}

	h.Close()
	if h.cb != nil {
		t.Error("callback registry retained after Close")
	}
	if h.cbHandle != 0 {
		t.Error("cgo.Handle not cleared after Close")
	}
}
