package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var glob = &AppSettings{
	Width:        1000,
	Height:       1000,
	MaxIter:      30,
	OrigoCap:     2.0,
	Fullscreen:   false,
	ScreenOffset: PositionF64{0, 0},
	Scale:        1.0,
}

func main() {
	a := app.New()

	wSett := a.NewWindow("GoLang Fractals glob")
	SettingsDisplay(wSett)

	wSett.SetOnClosed(func() {
		wView := a.NewWindow("GoLang Fractals Viewer")
		fmt.Printf("%+v\n", glob)

		// Set a key binding to close the window on Escape
		wView.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				wView.Close()
			}

			if key.Name == fyne.KeyF {
				glob.Fullscreen = !glob.Fullscreen
				wView.SetFullScreen(glob.Fullscreen)
			}

			step := 10.0
			if key.Name == fyne.KeyUp {
				glob.ScreenOffset.Y = glob.ScreenOffset.Y - step
				wView.Canvas().Refresh(wView.Content())

			}

			if key.Name == fyne.KeyDown {
				glob.ScreenOffset.Y = glob.ScreenOffset.Y + step
				wView.Canvas().Refresh(wView.Content())
			}

			if key.Name == fyne.KeyRight {
				glob.ScreenOffset.X = glob.ScreenOffset.X + step
				wView.Canvas().Refresh(wView.Content())
			}

			if key.Name == fyne.KeyLeft {
				glob.ScreenOffset.X = glob.ScreenOffset.X - step
				wView.Canvas().Refresh(wView.Content())
			}
		})

		ViewerDisplay(wView)
	})

	a.Run()
}
