package main

import (
	"embed"
	"fyne.io/fyne/v2/app"
	"github.com/lijingbo8119/minesweeper-fyne/core"
	_ "image/gif"
	_ "image/png"
)

//go:embed images/*
var images embed.FS

func main() {
	core.Init(images)

	a := app.New()
	a.Settings().SetTheme(&core.MyTheme{})

	w := a.NewWindow("Minesweeper")
	w.SetFixedSize(true)
	w.SetPadded(false)

	core.State.SetWindow(w)
	core.State.SetMatrixParamAndRender(9, 9, 10)

	w.ShowAndRun()
}
