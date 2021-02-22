package main

import (
	"embed"
	"fyne.io/fyne/v2/app"
	"fynetest/internal"
	"fynetest/myTheme"
	_ "image/gif"
	_ "image/png"
)

//go:embed images/*
var images embed.FS

func main() {
	internal.Init(images)

	a := app.New()
	a.Settings().SetTheme(&myTheme.MyTheme{})

	w := a.NewWindow("Minesweeper")
	w.SetFixedSize(true)
	w.SetPadded(false)

	internal.State.SetWindow(w)
	internal.State.SetMatrixParamAndRender(9, 9, 10)

	w.ShowAndRun()
}
