package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	settings := Settings{
		Width:      1000,
		Height:     1000,
		MaxIter:    30,
		OrigoCap:   2.0,
		Fullscreen: false,
	}

	wSett := a.NewWindow("GoLang Fractals Settings")
	SettingsDisplay(wSett, &settings)

	wSett.SetOnClosed(func() {
		wView := a.NewWindow("GoLang Fractals Viewer")
		fmt.Printf("%+v\n", settings)

		// Set a key binding to close the window on Escape
		wView.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				wView.Close()
			}

			if key.Name == fyne.KeyF {
				settings.Fullscreen = !settings.Fullscreen
				wView.SetFullScreen(settings.Fullscreen)
			}
		})

		ViewerDisplay(wView, settings)
	})

	a.Run()
}
