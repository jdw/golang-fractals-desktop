package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	settings := Settings{
		Width:        1000,
		Height:       1000,
		MaxIter:      30,
		OrigoCap:     2.0,
		Fullscreen:   false,
		ScreenOffset: PositionI64{0, 0},
		Scale:        1.0,
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

			step := int64(10)
			if key.Name == fyne.KeyUp {
				settings.ScreenOffset.Y = settings.ScreenOffset.Y - step
				wView.Canvas().Refresh(wView.Content())

			}

			if key.Name == fyne.KeyDown {
				settings.ScreenOffset.Y = settings.ScreenOffset.Y + step
				wView.Canvas().Refresh(wView.Content())
			}

			if key.Name == fyne.KeyRight {
				settings.ScreenOffset.X = settings.ScreenOffset.X + step
				wView.Canvas().Refresh(wView.Content())
			}

			if key.Name == fyne.KeyLeft {
				settings.ScreenOffset.X = settings.ScreenOffset.X - step
				wView.Canvas().Refresh(wView.Content())
			}
		})

		ViewerDisplay(wView, &settings)
	})

	a.Run()
}
