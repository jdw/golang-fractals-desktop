package main

type Controller struct {
	Model    *Model
	Settings Settings
}

func NewController(settings Settings) *Controller {
	return &Controller{

		Settings: settings,
	}
}

func (c *Controller) GetPixel(pos PositionI64, settings Settings) Pixel {
	if !c.Model.HasPixel(pos, settings) {
		posMandelbrot := Position.FromScreenCoordinateToMandelbrotCoordinates(int(pos.X), int(pos.Y), settings.Width, settings.Height)
		res := Mandelbrot.calculate(posMandelbrot, settings)
		return Pixel{res}
	}

	return Pixel{0}
}
