package loggerfile

import (
	"testing"
)

func TestShouldReturnWriter(t *testing.T) {
	result := New("path")

	if result == nil {
		t.Errorf("should return valid io.Writer, got: nil")
	}
}

func TestShouldReturnNilWithoutPath(t *testing.T) {
	result := New("")

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}
