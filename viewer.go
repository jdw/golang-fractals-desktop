package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ViewerDisplay(w fyne.Window, settings *Settings) {
	var view *View
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			if x == 0 && y == 0 {
				view = NewView(*settings)
			}

			res := view.GetScreenPixel(x, y).Iterations
			charge := float64(res) / float64(settings.MaxIter)

			ret := color.RGBA{R: uint8(charge * 255),
				G: uint8(charge * 255),
				B: uint8(charge * 255), A: 0xff}

			return ret
		})

	raster.Resize(fyne.NewSize(float32(settings.Width), float32(settings.Height)))
	raster.SetMinSize(fyne.NewSize(float32(settings.Width), float32(settings.Height)))
	w.SetContent(raster)

	if settings.Fullscreen {
		w.SetFullScreen(true)
	} else {
		w.SetFullScreen(false)
		//w.Resize(fyne.NewSize(float32(settings.Width), float32(settings.Height)))
		w.CenterOnScreen()
	}

	w.Show()
}
