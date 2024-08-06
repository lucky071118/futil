package cmd

import (
	"testing"
)

func TestCountLine(t *testing.T) {
	result, _ := CountLine("testdata/myfile.txt")

	expected := 4

	if result != expected {
		t.Errorf("GetLineCount('testdata/myfile.txt') = %d; want %d", result, expected)
	}
}

func TestCountLineWithFileNotFound(t *testing.T) {
	_, err := CountLine("testdata/file_not_found.txt")

	if err == nil {
		t.Errorf("GetLineCount('testdata/file_not_found.txt') = _, nil; want _, error")
	}
}
