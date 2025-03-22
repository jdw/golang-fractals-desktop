package main

import (
	"fmt"

	"fyne.io/fyne/v2"
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
		// Set a key binding to close the window on Escape
		wView.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				wView.Close()
			}
		})
		fmt.Printf("%+v\n", settings)
		ViewerDisplay(wView, settings)
	})

	a.Run()
}
