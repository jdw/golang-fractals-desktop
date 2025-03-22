package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	settings := Settings{
		Width:   1920,
		Height:  1080,
		MaxIter: 30,
	}

	wSett := a.NewWindow("GoLang Fractals Settings")
	SettingsDisplay(wSett, &settings)

	wSett.SetOnClosed(func() {
		wView := a.NewWindow("GoLang Fractals Viewer")
		fmt.Printf("%+v\n", settings)
		ViewerDisplay(wView, settings)
	})

	a.Run()
}
