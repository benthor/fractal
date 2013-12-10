package main

import (
	//"math/cmplx"
	"image/color"
	"image/png"
	"os"
	"fmt"
	"./fractal"
)


func mandelbrotRecur(z, c complex128, max float64, depth int) color.Color {
	if depth > 1 {
		a, b := real(z), imag(z)
		znew := complex((a*a-b*b),2*a*b) + c
		/*if cmplx.Abs(znew) < max {
			return mandelbrot(znew, c, max, depth-1)
		}*/
		if a*a+b*b < max {
			return mandelbrotRecur(znew, c, max, depth-1)
		}
	}
	fmt.Printf("bop")
	return color.RGBA{uint8(255/depth), uint8(255/depth), uint8(255/depth), 255}
}



func mandelbrotIter(z, c complex128, max float64, depth int, r []complex128) color.Color {
	var a, b float64
	iteration := 0
	l := len(r)
iter:
	for a*a+b*b < max && iteration < depth{
		iteration++
		a, b = real(z), imag(z)
		z = complex((a*a-b*b),2*a*b) + c
		if iteration % l == 0 {
			for i:=l-1; i>0; i-- {
				if r[i] == z {
					iteration = depth
					break iter 
				}
			}
		}
		r[iteration%l] = z

	}
	return color.RGBA{uint8(255*iteration/depth), uint8((255*iteration)/depth), uint8(255*iteration/depth), 255}
}

func loopdetect(path []complex128, c complex128) {
	
}


func main() {
	f := fractal.NewFractal(complex(-3, -1.5), complex(1.5, 1.5), 1000)
	d := 1000
	r := make([]complex128, 40)
	f.Apply2All(
		func(c complex128) color.Color {
			return mandelbrotIter(complex(0,0), c, 4, d, r)
		})

	fo, err := os.Create("output.png")
	defer fo.Close()

	if err != nil {
		panic(err)
	}
	png.Encode(fo, f)
}