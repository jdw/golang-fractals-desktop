package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ViewerDisplay(w fyne.Window, settings Settings) {
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			fY := ((float64(y) / float64(h)) - 0.5) * 5.0
			fX := ((float64(x) / float64(w)) - 0.5) * 5.0

			if x == 0 && y == 0 {
				println(x, y, fX, fY)
			}

			if x == 0 && y == h-1 {
				println(x, y, fX, fY)
			}

			if x == w-1 && y == 0 {
				println(x, y, fX, fY)
			}

			if x == w-1 && y == h-1 {
				println(x, y, fX, fY)
			}

			res := Mandelbrot.calculate(fX, fY, settings)
			charge := float64(res) / float64(settings.MaxIter)

			ret := color.RGBA{R: uint8(charge * 255),
				G: uint8(charge * 255),
				B: uint8(charge * 255), A: 0xff}

			return ret
		})

	raster.Resize(fyne.NewSize(float32(settings.Width), float32(settings.Height)))
	//raster.FillMode = canvas.ImageFillContain
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
