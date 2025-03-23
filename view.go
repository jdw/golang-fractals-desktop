package main

type ViewSettings struct {
	Offset PositionF64 // Offset is the position of the viewport in the world
	Size   PositionF64 // Size is the size of the viewport in pixels
	Scale  float64     // Scale is the zoom level of the view onto the model
}

func NewViewSettings(settings *AppSettings) *ViewSettings {
	return &ViewSettings{
		Offset: PositionF64{X: settings.ScreenOffset.X, Y: settings.ScreenOffset.Y},
		Size:   PositionF64{X: float64(settings.Width), Y: float64(settings.Height)},
		Scale:  settings.Scale,
	}
}

type View struct {
	Controller *Controller
	Settings   *ViewSettings // Settings contains the configuration for the specific frame
}

func NewView(settings *AppSettings) *View {
	return &View{
		Controller: NewController(settings),
		Settings:   NewViewSettings(settings),
	}
}

func (v *View) GetScreenPixel(pos PositionF64) Pixel {
	tPos := PositionF64{X: pos.X - v.Settings.Offset.X, Y: pos.Y - v.Settings.Offset.Y}
	return v.Controller.GetPixel(tPos, glob)
}
