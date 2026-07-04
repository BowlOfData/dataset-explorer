package editor

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const appTitle = "Text Editor"

// Editor is a simple single-document text editor window.
type Editor struct {
	app   fyne.App
	win   fyne.Window
	entry *widget.Entry

	filePath  string // "" means the document has never been saved to a path
	savedText string // document content as of the last successful load/save
}

// New constructs an Editor bound to the given Fyne application.
func New(a fyne.App) *Editor {
	e := &Editor{
		app:   a,
		entry: widget.NewMultiLineEntry(),
	}
	e.entry.Wrapping = fyne.TextWrapWord
	e.entry.TextStyle = fyne.TextStyle{Monospace: true}
	e.entry.OnChanged = func(string) { e.updateTitle() }
	return e
}

// Run builds the window and starts the Fyne event loop. Blocks until the app quits.
func (e *Editor) Run() {
	e.win = e.app.NewWindow(appTitle)
	e.win.SetMaster()
	e.win.Resize(fyne.NewSize(800, 600))
	e.win.SetContent(container.NewPadded(e.entry))
	e.win.SetMainMenu(buildMainMenu(e))
	e.win.SetCloseIntercept(func() {
		e.confirmDiscardIfDirty(func() { e.win.Close() })
	})
	e.updateTitle()
	e.win.ShowAndRun()
}

// isDirty reports whether the document has unsaved changes.
func (e *Editor) isDirty() bool {
	return e.entry.Text != e.savedText
}

func displayName(path string) string {
	if path == "" {
		return "Untitled"
	}
	return path
}

// formatTitle builds the window title from the current path and dirty state.
// Pure function, kept separate from Editor state so it's trivially unit-testable.
func formatTitle(path string, dirty bool) string {
	name := displayName(path)
	if dirty {
		return fmt.Sprintf("*%s — %s", name, appTitle)
	}
	return fmt.Sprintf("%s — %s", name, appTitle)
}

func (e *Editor) updateTitle() {
	if e.win == nil {
		return
	}
	e.win.SetTitle(formatTitle(e.filePath, e.isDirty()))
}

// confirmDiscardIfDirty runs next immediately if there are no unsaved changes,
// otherwise it asks the user to confirm discarding them first.
func (e *Editor) confirmDiscardIfDirty(next func()) {
	if !e.isDirty() {
		next()
		return
	}
	dialog.NewConfirm(
		"Unsaved changes",
		"You have unsaved changes. Discard them?",
		func(discard bool) {
			if discard {
				next()
			}
		},
		e.win,
	).Show()
}
