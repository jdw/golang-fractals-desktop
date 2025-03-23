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

func (c *Controller) GetPixel(pos PositionI64, settings *AppSettings) Pixel {
	if !c.Model.HasPixel(pos, settings) {
		posMandelbrot := Position.FromScreenCoordinateToMandelbrotCoordinates(int(pos.X), int(pos.Y), settings)
		res := Mandelbrot.calculate(posMandelbrot, settings)

		return Pixel{res}
	}

	return Pixel{0}
}
