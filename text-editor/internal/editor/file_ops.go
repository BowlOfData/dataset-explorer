package editor

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// doNew clears the document, prompting to discard unsaved changes first.
func (e *Editor) doNew() {
	e.confirmDiscardIfDirty(func() {
		e.entry.SetText("")
		e.filePath = ""
		e.savedText = ""
		e.updateTitle()
	})
}

// doOpen shows a file-open dialog and loads the chosen file.
func (e *Editor) doOpen() {
	e.confirmDiscardIfDirty(func() {
		d := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, e.win)
				return
			}
			if r == nil {
				return // user cancelled
			}
			path := r.URI().Path()
			r.Close()
			if err := e.loadFile(path); err != nil {
				dialog.ShowError(err, e.win)
			}
		}, e.win)
		d.Show()
	})
}

// doSave writes the document to its current path, or delegates to Save As
// if the document has no path yet.
func (e *Editor) doSave() {
	if e.filePath == "" {
		e.doSaveAs()
		return
	}
	if err := e.writeFile(e.filePath); err != nil {
		dialog.ShowError(err, e.win)
	}
}

// doSaveAs shows a file-save dialog and writes the document to the chosen path.
// The dialog callback is asynchronous, so this cannot report success synchronously.
func (e *Editor) doSaveAs() {
	d := dialog.NewFileSave(func(w fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, e.win)
			return
		}
		if w == nil {
			return // user cancelled
		}
		path := w.URI().Path()
		w.Close()
		if err := e.writeFile(path); err != nil {
			dialog.ShowError(err, e.win)
		}
	}, e.win)
	d.Show()
}

// loadFile reads path from disk and replaces the current document with its contents.
func (e *Editor) loadFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(data)
	e.entry.SetText(text) // intentionally resets Entry's undo history for the new document
	e.filePath = path
	e.savedText = text
	e.updateTitle()
	return nil
}

// writeFile writes the current document contents to path.
func (e *Editor) writeFile(path string) error {
	if err := os.WriteFile(path, []byte(e.entry.Text), 0644); err != nil {
		return err
	}
	e.filePath = path
	e.savedText = e.entry.Text
	e.updateTitle()
	return nil
}
