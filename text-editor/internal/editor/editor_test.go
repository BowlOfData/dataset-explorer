package editor

import (
	"os"
	"path/filepath"
	"testing"

	"fyne.io/fyne/v2/test"
)

func newTestEditor() *Editor {
	return New(test.NewApp())
}

func TestFormatTitle(t *testing.T) {
	cases := []struct {
		path  string
		dirty bool
		want  string
	}{
		{"", false, "Untitled — Text Editor"},
		{"", true, "*Untitled — Text Editor"},
		{"/tmp/notes.txt", false, "/tmp/notes.txt — Text Editor"},
		{"/tmp/notes.txt", true, "*/tmp/notes.txt — Text Editor"},
	}
	for _, c := range cases {
		if got := formatTitle(c.path, c.dirty); got != c.want {
			t.Errorf("formatTitle(%q, %v) = %q, want %q", c.path, c.dirty, got, c.want)
		}
	}
}

func TestIsDirty(t *testing.T) {
	e := newTestEditor()

	if e.isDirty() {
		t.Fatal("new editor should not be dirty")
	}

	e.entry.SetText("hello")
	if !e.isDirty() {
		t.Fatal("editor should be dirty after text change")
	}

	e.savedText = "hello"
	if e.isDirty() {
		t.Fatal("editor should not be dirty once savedText matches entry text")
	}
}

func TestLoadFileRoundTrip(t *testing.T) {
	e := newTestEditor()
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	want := "line one\nline two\n"
	if err := os.WriteFile(path, []byte(want), 0644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	if err := e.loadFile(path); err != nil {
		t.Fatalf("loadFile: %v", err)
	}
	if e.entry.Text != want {
		t.Errorf("entry.Text = %q, want %q", e.entry.Text, want)
	}
	if e.filePath != path {
		t.Errorf("filePath = %q, want %q", e.filePath, path)
	}
	if e.isDirty() {
		t.Error("editor should not be dirty right after loading a file")
	}
}

func TestWriteFileRoundTrip(t *testing.T) {
	e := newTestEditor()
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	e.entry.SetText("some content")
	if err := e.writeFile(path); err != nil {
		t.Fatalf("writeFile: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading back written file: %v", err)
	}
	if string(got) != "some content" {
		t.Errorf("written content = %q, want %q", string(got), "some content")
	}
	if e.isDirty() {
		t.Error("editor should not be dirty right after saving")
	}
	if e.filePath != path {
		t.Errorf("filePath = %q, want %q", e.filePath, path)
	}
}

func TestDoNewClearsDocument(t *testing.T) {
	e := newTestEditor()
	e.win = test.NewWindow(e.entry)
	defer e.win.Close()

	e.entry.SetText("something")
	e.filePath = "/tmp/whatever.txt"
	e.savedText = "something" // not dirty, so doNew proceeds without a confirm dialog

	e.doNew()

	if e.entry.Text != "" {
		t.Errorf("entry.Text = %q, want empty after New", e.entry.Text)
	}
	if e.filePath != "" {
		t.Errorf("filePath = %q, want empty after New", e.filePath)
	}
}
