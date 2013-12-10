package fractal

import (
	"image"
	"image/color"
)

type Fractal struct {
	LowerLeft complex128
	UpperRight complex128
	Scale float64
	RGBA image.RGBA
}

func NewFractal(lowerLeft, upperRight complex128, scale float64) Fractal {
	w := int((real(upperRight) - real(lowerLeft))*scale)
	h := int((imag(upperRight) - imag(lowerLeft))*scale)
	return Fractal{lowerLeft, upperRight, scale, *image.NewRGBA(image.Rect(0,0,w,h))}
}


func (f Fractal) Bounds() image.Rectangle{
	return f.RGBA.Bounds()
}

func (f Fractal) ColorModel() color.Model {
	return f.RGBA.ColorModel()
}

func (f Fractal) At (x, y int) color.Color {
	return f.RGBA.At(x,y)
}

func (f Fractal) Set(x, y int, c color.Color) {
	f.RGBA.Set(x,y,c)
}

func (f Fractal) Complex2Point(c complex128) image.Point {
	rec := f.Bounds()
	x := int((real(c) + real(f.LowerLeft))*f.Scale) + rec.Min.X
	y := int((imag(c) + imag(f.LowerLeft))*f.Scale) + rec.Min.Y
	return image.Pt(x,y)
}


func (f Fractal) Point2Complex(p image.Point) complex128 {
	rec := f.Bounds()
	x := float64(p.X - rec.Min.X)/f.Scale + real(f.LowerLeft)
	y := float64(p.Y - rec.Min.Y)/f.Scale + imag(f.LowerLeft)
	return complex(x,y)
}

func (f Fractal) Apply2All(pos2color func(complex128) color.Color) {
	rec := f.Bounds()
	for y := rec.Min.Y; y < rec.Max.Y; y++ {
		for x := rec.Min.X; x < rec.Max.X; x++ {
			f.Set(x,y,pos2color(f.Point2Complex(image.Pt(x,y))))
		}
	}
}