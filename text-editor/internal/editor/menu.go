package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func shortcut(key fyne.KeyName) *desktop.CustomShortcut {
	return &desktop.CustomShortcut{KeyName: key, Modifier: fyne.KeyModifierShortcutDefault}
}

// buildMainMenu constructs the File/Edit menu bar for the editor window.
// Keyboard shortcuts use fyne.KeyModifierShortcutDefault, which resolves to
// Ctrl on Linux/Windows and Cmd on macOS from a single code path.
func buildMainMenu(e *Editor) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", e.doNew)
	newItem.Shortcut = shortcut(fyne.KeyN)

	openItem := fyne.NewMenuItem("Open…", e.doOpen)
	openItem.Shortcut = shortcut(fyne.KeyO)

	saveItem := fyne.NewMenuItem("Save", e.doSave)
	saveItem.Shortcut = shortcut(fyne.KeyS)

	saveAsItem := fyne.NewMenuItem("Save As…", e.doSaveAs)
	saveAsItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyS,
		Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierShift,
	}

	quitItem := fyne.NewMenuItem("Quit", func() {
		e.confirmDiscardIfDirty(e.app.Quit)
	})
	quitItem.Shortcut = shortcut(fyne.KeyQ)
	quitItem.IsQuit = true

	fileMenu := fyne.NewMenu("File",
		newItem,
		openItem,
		fyne.NewMenuItemSeparator(),
		saveItem,
		saveAsItem,
		fyne.NewMenuItemSeparator(),
		quitItem,
	)

	undoItem := fyne.NewMenuItem("Undo", e.entry.Undo)
	undoItem.Shortcut = shortcut(fyne.KeyZ)

	redoItem := fyne.NewMenuItem("Redo", e.entry.Redo)
	redoItem.Shortcut = shortcut(fyne.KeyY)

	editMenu := fyne.NewMenu("Edit", undoItem, redoItem)

	return fyne.NewMainMenu(fileMenu, editMenu)
}
