package poker

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, []Player{{Name: "A", Wins: 10}})
	defer clean()

	want := "abc"
	tape := &tape{file}
	tape.Write([]byte(want))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
