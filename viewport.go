package main

type Viewport struct {
	data   map[Position]Pixel
	offset Position // Offset is the position of the viewport in the world
	scale  float64  // Scale is the zoom level of the viewport
}

func NewViewport() *Viewport {
	return &Viewport{
		data:   make(map[Position]Pixel),
		offset: Position{X: 0, Y: 0},
		scale:  1.0,
	}
}

func (v *Viewport) GetPixel(p Position) Pixel {
	return v.data[p]
}

func (v *Viewport) SetPixel(p Position, pixel Pixel) {
	v.data[p] = pixel
}

func (v *Viewport) CalculatePixels(topLeft, bottomRight Position, settings Settings) {
	for x := topLeft.X; x <= bottomRight.X; x++ {
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			iterations := Mandelbrot.calculate(x, y, settings)
			pxl := Pixel{iterations: iterations, maxIterations: settings.MaxIter, origoCap: settings.OrigoCap}
			v.SetPixel(Position{X: x, Y: y}, pxl)
		}
	}
}

func (v *Viewport) SetOffset(x, y int64) {
	v.offset.X = x
	v.offset.Y = y
}

func (v *Viewport) SetScale(scale float64) {
	v.scale = scale
}

func (v *Viewport) TranslateScreenToViewportPosition(screen, size Position) Position {
	return screen
}

func (v *Viewport) IsEmpty() bool {
	return len(v.data) == 0
}
