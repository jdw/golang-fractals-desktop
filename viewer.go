package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ViewerDisplay(w fyne.Window, settings Settings) {
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			fY := float64(y) / float64(h)
			fX := float64(x) / float64(w)
			//widthScalar := float64(settings.Width) / float64(settings.Height)
			//heightScalar := float64(settings.Height) / float64(settings.Width)

			res := FractalUtil.calculate((fX-0.5)*5.0, (fY-0.5)*5.0, settings.MaxIter)
			charge := float64(res) / float64(settings.MaxIter)

			ret := color.RGBA{R: uint8(charge * 255),
				G: uint8(charge * 255),
				B: uint8(charge * 255), A: 0xff}

			return ret
		})

	w.SetContent(raster)
	if settings.Fullscreen {
		w.SetFullScreen(true)
	} else {
		w.SetFullScreen(false)
		w.Resize(fyne.NewSize(float32(settings.Width), float32(settings.Height)))
		w.CenterOnScreen()
	}

	w.Show()
}
