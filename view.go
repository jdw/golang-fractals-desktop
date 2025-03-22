package main

type View struct {
	data       map[PositionI64]Pixel
	offset     PositionI64 // Offset is the position of the viewport in the world
	size       PositionI64 // Size is the size of the viewport in pixels
	scale      float64     // Scale is the zoom level of the viewport
	Controller *Controller
	Settings   Settings // Settings contains the configuration for the view
}

func NewView(settings Settings) *View {
	return &View{
		data:       make(map[PositionI64]Pixel),
		offset:     PositionI64{X: 0, Y: 0},
		size:       PositionI64{X: int64(settings.Width), Y: int64(settings.Height)},
		scale:      1.0,
		Controller: NewController(settings),
		Settings:   settings,
	}
}

func (v *View) GetScreenPixel(x, y int) Pixel {
	return v.Controller.GetPixel(PositionI64{X: int64(x), Y: int64(y)}, v.Settings)
}
