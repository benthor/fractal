package main

import (
	//"math/cmplx"
	"image/color"
	"image/png"
	"os"

	"./fractal"
)

func mandelbrot(z, c complex128, max float64, depth int) color.Color {
	if depth > 1 {
		a, b := real(z), imag(z)
		znew := complex((a*a-b*b),2*a*b) + c
		/*if cmplx.Abs(znew) < max {
			return mandelbrot(znew, c, max, depth-1)
		}*/
		if a*a+b*b < max {
			return mandelbrot(znew, c, max, depth-1)
		}


	}
	return color.RGBA{uint8(255/depth), uint8(255/depth), uint8(255/depth), 255}
}

func mandelbrotIter(z, c complex128, max float64, depth int) color.Color {
	var a, b float64
	iteration := depth
	for a*a+b*b < max && iteration > 0{
		a, b = real(z), imag(z)
		z = complex((a*a-b*b),2*a*b) + c
		iteration--
	}
	return color.RGBA{uint8(255*iteration/depth), uint8((255*iteration)/depth), uint8(255*iteration/depth), 255}
}


func main() {
	f := fractal.NewFractal(complex(-3, -1.5), complex(1.5, 1.5), 1000)
	f.Apply2All(
		func(c complex128) color.Color {
			return mandelbrot(complex(0,0), c, 4, 100)
		})

	fo, err := os.Create("output.png")
	defer fo.Close()

	if err != nil {
		panic(err)
	}
	png.Encode(fo, f)
}