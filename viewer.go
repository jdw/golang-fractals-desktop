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
			if x == 0 && y == 0 && view == nil {
				view = NewView(glob)
			}

			if x == 0 && y == 0 {
				oldSettings := view.Settings
				view.Settings = NewViewSettings(glob)

				if oldSettings.Scale != view.Settings.Scale {
					view.TransferTexture()
				}
			}

			return view.GetScreenPixel(x, y)
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
