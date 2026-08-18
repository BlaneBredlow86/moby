package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopierFlushesPartialLine(t *testing.T) {
	input := "final-partial-line"
	src := strings.NewReader(input)
	dst := &bytes.Buffer{}
	copier := &Copier{}

	err := copier.Copy(src, dst)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	if dst.String() != input {
		t.Errorf("Expected %q, got %q", input, dst.String())
	}
}