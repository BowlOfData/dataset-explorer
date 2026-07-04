package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/bowlofdata/text-editor/internal/editor"
)

func main() {
	a := app.New()
	e := editor.New(a)
	e.Run()
}
