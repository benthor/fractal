// Package fractal defines functions to compute fractals
package fractal

import (
	"image"
	//"fmt"
	"math/cmplx"
	"image/color"
)

func MandelbrotIteration(z complex128, c complex128) complex128 {
	return cmplx.Pow(z,2) + c
}

func Bounded(function func(complex128, complex128) complex128, z complex128, c complex128, abs float64, depth int) int {
	if depth > 0 {
		znew := function(z, c)
		if cmplx.Abs(znew) < abs {
			r := Bounded(function, znew, c, abs, depth-1)
			//fmt.Println(real(c),imag(c), real(znew), imag(znew), cmplx.Abs(znew), r)
			return r
		}
	}
	return depth
}

func Scale(inlower, inupper complex128, rec image.Rectangle) (xfactor, yfactor float64) {
	xfactor = float64(rec.Max.X - rec.Min.X) / (real(inupper) - real(inlower))
	yfactor = float64(rec.Max.Y - rec.Min.Y) / (imag(inupper) - imag(inlower))
	return
}

func Complex2Point(inlower, inupper complex128, rec image.Rectangle) func(complex128) image.Point {
	xfactor, yfactor := Scale(inlower, inupper, rec)
	return func(cp complex128) image.Point {
		// FIXME, doesn't take into account lower corner
		return image.Pt(rec.Min.X + int(real(cp)*xfactor), rec.Min.Y + int(imag(cp)*yfactor))
	}
}

func Point2Complex(inlower, inupper complex128, rec image.Rectangle) func(image.Point) complex128 {
	xfactor, yfactor := Scale(inlower, inupper, rec)
	return func(p image.Point) complex128 {
		return complex(float64(p.X - rec.Min.X)/xfactor + real(inlower), float64(p.Y-rec.Min.Y)/yfactor + imag(inlower))
	}
}


func Iter(img image.RGBA, Do func(x,y int) color.Color) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			img.Set(x, y, Do(x,y))
		}
		//fmt.Println()
	}
}