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

func (p *PositionHidden) TransformViewScalarToMandelbrotCoordinate(pos PositionF64, settings *AppSettings) PositionF64 { // TODO Bad namn
	scalar := 2.5

	return PositionF64{X: pos.X * scalar, Y: pos.Y * scalar}
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

func (p *PositionHidden) TransformScreenPositionsToViewScalar(x, y int, settings *AppSettings) PositionF64 {
	fX := float64(x)
	fY := float64(y)
	w := settings.Width
	h := settings.Height
	tX := fX / w
	tY := fY / h

	return PositionF64{X: tX, Y: tY}
}

func (p *PositionHidden) TransformViewScalarToModelCoordinate(pos PositionF64, settings *AppSettings) PositionI64 {
	//pixelQuota := float64(settings.ModelDimensions.X) / settings.Width
	tX := pos.X * float64(settings.ModelDimensions.X)
	tY := pos.Y * float64(settings.ModelDimensions.Y)

	return PositionI64{X: int64(tX), Y: int64(tY)}
}

func (p *PositionHidden) TransformScreenPositionsToViewOrigoCenteredScalar(x, y int, settings *AppSettings) PositionF64 {
	w := settings.Width
	h := settings.Height
	scalarX := float64(x) / float64(w)
	scalarY := float64(y) / float64(h)

	// Optional: Scale to a specific range (e.g., -1 to 1)
	scalarX = (scalarX * 2) - 1
	scalarY = (scalarY * 2) - 1

	return PositionF64{X: scalarX, Y: scalarY}
}

func (p *PositionI64) GetPositionF64() PositionF64 {
	return PositionF64{X: float64(p.X), Y: float64(p.Y)}
}
