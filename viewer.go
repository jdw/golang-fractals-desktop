package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ViewerDisplay(w fyne.Window, settings Settings) {
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			fY := float64(y) / float64(h)
			fX := float64(x) / float64(w)
			res := FractalUtil.calculate(fX, fY, settings.MaxIter)

			ret := color.RGBA{R: uint8(56),
				G: uint8(56),
				B: uint8(56), A: 0xff}

			if res == 10 {
				ret = color.RGBA{R: uint8(rand.Intn(255)),
					G: uint8(rand.Intn(255)),
					B: uint8(rand.Intn(255)), A: 0xff}
			}
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
