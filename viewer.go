package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ViewerDisplay(w fyne.Window) {
	var view *View
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			if x == 0 && y == 0 {
				view = NewView(glob)
			}

			res := view.GetScreenPixel(x, y).Iterations
			charge := float64(res) / float64(glob.MaxIter)

			ret := color.RGBA{R: uint8(charge * 255),
				G: uint8(charge * 255),
				B: uint8(charge * 255), A: 0xff}

			return ret
		})

	raster.Resize(fyne.NewSize(float32(glob.Width), float32(glob.Height)))
	raster.SetMinSize(fyne.NewSize(float32(glob.Width), float32(glob.Height)))
	w.SetContent(raster)

	if glob.Fullscreen {
		w.SetFullScreen(true)
	} else {
		w.SetFullScreen(false)
		w.CenterOnScreen()
	}

	w.Show()
}
