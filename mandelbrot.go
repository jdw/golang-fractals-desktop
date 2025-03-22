package main

import (
	"math/cmplx"
)

type MandelbrotHidden struct {
}

var Mandelbrot MandelbrotHidden

func (f *MandelbrotHidden) calculate(x, y float64, settings Settings) int {
	ret := 0
	z := complex(0, 0)
	c := complex(x, y)

	for cmplx.Abs(z) < settings.OrigoCap && ret < settings.MaxIter {
		z = z*z + c
		ret++
	}

	return ret
}

// func (f *MandelbrotHidden) calculate(x, y int64, settings Settings) int {
// 	ret := 0
// 	z := complex(0, 0)
// 	// Scale the coordinates to the range of the complex plane
// 	// The Mandelbrot set is defined for the complex plane, so we need to scale the coordinates
// 	// to the range of the complex plane.
// 	// The range of the complex plane is from -2 to 2 in both the real and imaginary parts
// 	// We can scale the coordinates to the range of the complex plane by dividing by the maximum
// 	// value of the coordinates and multiplying by the range of the complex plane
// 	// The maximum value of the coordinates is the maximum of the width and height of the image
// 	// The range of the complex plane is from -2 to 2 in both the real and imaginary parts
// 	// The range of the complex plane is 4 in both the real and imaginary parts
// 	var qX, qY float64
// 	if x == 0 {
// 		qX = 0.0
// 	} else if x > 0 {
// 		qX = math.Abs(float64(x)) / math.MaxFloat64
// 	} else {
// 		qX = -math.Abs(float64(x)) / math.MaxFloat64
// 	}

// 	if y == 0 {
// 		qY = 0
// 	} else if y > 0 {
// 		qY = math.Abs(float64(y)) / math.MaxFloat64
// 	} else {
// 		qY = -math.Abs(float64(y)) / math.MaxFloat64
// 	}

// 	qX *= 2.0
// 	qY *= 2.0

// 	c := complex(qX, qY)

// 	fmt.Println("X:", x, "Y:", y, "QX:", qX, "QY:", qY, "C:", c)

// 	for cmplx.Abs(z) < settings.OrigoCap && ret < settings.MaxIter {
// 		z = z*z + c
// 		ret++
// 	}

// 	return ret
// }
