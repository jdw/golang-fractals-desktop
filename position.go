package main

type PositionI64 struct {
	X int64
	Y int64
}

type PositionF64 struct {
	X float64
	Y float64
}

type PositionHidden struct {
}

var Position PositionHidden

func (p *PositionHidden) TransformScreenCoordinateToMandelbrotCoordinate(x, y int, settings *AppSettings) PositionF64 {
	w := settings.Width
	h := settings.Height
	fX := (float64(x)/w - 0.5) * 5.0
	fY := (float64(y)/h - 0.5) * 5.0

	return PositionF64{X: fX, Y: fY}
}

func (p *PositionHidden) TransformViewCoordinateToModelCoordinate(x, y, w, h int) PositionI64 {
	wX := (float64(x) / float64(w)) * glob.ModelDimensions.GetPositionF64().X
	wY := (float64(y) / float64(h)) * glob.ModelDimensions.GetPositionF64().Y

	return PositionI64{X: int64(wX), Y: int64(wY)}
}

func (p *PositionI64) GetPositionF64() PositionF64 {
	// USe math.Float64frombits() ????
	return PositionF64{X: float64(p.X), Y: float64(p.Y)}
}
