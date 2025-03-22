package main

import (
	"math/cmplx"
)

type MandelbrotHidden struct {
}

var Mandelbrot MandelbrotHidden

func (f *MandelbrotHidden) calculate(pos PositionF64, settings Settings) uint8 {
	ret := uint8(0)
	maxIter := uint8(settings.MaxIter)

	z := complex(0, 0)
	c := complex(pos.X, pos.Y)

	for cmplx.Abs(z) < settings.OrigoCap && ret < maxIter {
		z = z*z + c
		ret++
	}

	return ret
}
