package main

type ViewSettings struct {
	Offset      PositionF64 // Offset is the position of the viewport in the world
	TextureSize PositionF64 // Size is the size of the texture in pixels
	Scale       float64     // Scale is the zoom level of the view onto the model
}

func NewViewSettings(settings *AppSettings) *ViewSettings {
	return &ViewSettings{
		Offset:      PositionF64{X: settings.ScreenOffset.X, Y: settings.ScreenOffset.Y},
		TextureSize: PositionF64{X: settings.Width * 2, Y: settings.Height * 2},
		Scale:       settings.Scale,
	}
}

type View struct {
	Controller *Controller
	Settings   *ViewSettings // Settings contains the configuration for the specific frame
	Texture    map[PositionI64]Pixel
}

func NewView(settings *AppSettings) *View {
	ret := View{
		Controller: NewController(settings),
		Settings:   NewViewSettings(settings),
		Texture:    make(map[PositionI64]Pixel),
	}

	ret.TransferTexture()

	return &ret
}

func (v *View) GetScreenPixel(x, y int) Pixel {
	offsetX := x - int(v.Settings.Offset.X)
	offsetY := y - int(v.Settings.Offset.Y)

	return v.Texture[PositionI64{int64(offsetX), int64(offsetY)}]
}

func (v *View) TransferTexture() {
	textureSizeX := int(v.Settings.TextureSize.X)
	textureSizeY := int(v.Settings.TextureSize.Y)

	for x := 0; x < textureSizeX; x++ {
		for y := 0; y < textureSizeY; y++ {
			origoScalar := Position.TransformScreenPositionsToViewOrigoCenteredScalar(x, y, glob)
			pos := PositionF64{
				X: origoScalar.X * v.Settings.Scale,
				Y: origoScalar.Y * v.Settings.Scale,
			}

			v.Texture[PositionI64{int64(x), int64(y)}] = v.Controller.GetViewPixel(pos, glob)
		}
	}
}
