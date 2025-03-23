package main

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var glob = &AppSettings{
	Width:           1000,
	Height:          1000,
	MaxIter:         30,
	OrigoCap:        2.0,
	Fullscreen:      false,
	ScreenOffset:    PositionF64{0, 0},
	Scale:           1.0,
	ModelDimensions: PositionI64{math.MaxInt32, math.MaxInt32},
}

func main() {
	a := app.New()

	tUpperLeft := Position.TransformViewScalarToModelCoordinate(PositionF64{0, 0}, glob)
	tLowerRight := Position.TransformViewScalarToModelCoordinate(PositionF64{1.0, 1.0}, glob)
	fmt.Println("Upper left: ", tUpperLeft)
	fmt.Println("Lower right: ", tLowerRight)
	fmt.Println("Model dimensions: ", glob.ModelDimensions)
	wSett := a.NewWindow("GoLang Fractals Settings")
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

			if key.Name == fyne.KeyS {
				glob.Scale = glob.Scale * 1.1
				wView.Canvas().Refresh(wView.Content())
			}

			if key.Name == fyne.KeyA {
				glob.Scale = glob.Scale * 0.9
				wView.Canvas().Refresh(wView.Content())
			}

			fmt.Println(glob)
		})

		ViewerDisplay(wView)
	})

	a.Run()
}
