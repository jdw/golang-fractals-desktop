package main

type Controller struct {
	Model    *Model
	Settings *AppSettings
}

var model = NewModel()

func NewController(settings *AppSettings) *Controller {
	return &Controller{
		Model:    model,
		Settings: settings,
	}
}

func (c *Controller) GetViewPixel(pos PositionF64, settings *AppSettings) Pixel {
	modelPos := Position.TransformViewScalarToModelCoordinate(pos, settings)
	if !c.Model.HasPixel(modelPos, settings) {
		tPos := Position.TransformViewScalarToMandelbrotCoordinate(pos, settings)
		res := Mandelbrot.calculate(tPos, settings)

		return Pixel{res}
	}

	return Pixel{0}
}
