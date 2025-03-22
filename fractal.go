package main

import (
	"math/cmplx"
)

type Fractal struct {
}

var FractalUtil Fractal

func (f *Fractal) calculate(x, y float64, maxIterations int) int {
	ret := 0
	z := complex(0, 0)
	c := complex(x, y)

	for cmplx.Abs(z) < 2 && ret < maxIterations {
		z = z*z + c
		ret++
	}

	return ret
}
