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

func (p *PositionHidden) TransformScreenCoordinateToMandelbrotCoordinate(pos PositionF64, settings *AppSettings) PositionF64 {
	x := pos.X
	y := pos.Y
	w := settings.Width
	h := settings.Height
	tX := (x/w - 0.5) * 5.0
	tY := (y/h - 0.5) * 5.0

	return PositionF64{X: tX, Y: tY}
}

func (p *PositionHidden) TransformViewCoordinateToModelCoordinate(pos PositionF64, settings *AppSettings) PositionI64 {
	x := pos.X
	y := pos.Y
	w := settings.Width
	h := settings.Height
	tX := (x / w) * glob.ModelDimensions.GetPositionF64().X
	tY := (y / h) * glob.ModelDimensions.GetPositionF64().Y

	return PositionI64{X: int64(tX), Y: int64(tY)}
}

func (p *PositionI64) GetPositionF64() PositionF64 {
	return PositionF64{X: float64(p.X), Y: float64(p.Y)}
}
