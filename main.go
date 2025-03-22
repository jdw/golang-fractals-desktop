package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("GoLang Fractals Desktop")

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			fY := float64(y) / float64(h)
			fX := float64(x) / float64(w)
			res := FractalUtil.calculate(fX, fY, 30)

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
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(1920, 1080))
	w.ShowAndRun()
}
