package fractal

import (
	"testing"
	"image"
	"image/color"
	"image/png"
	//"io"
	//"fmt"
	"os"
)

func TestManelbrot(t *testing.T) {
	m := image.NewRGBA(image.Rect(0,0,1800,1200))
	lower := complex(-2, -1)*1.5
	upper := complex(1, 1)*1.5
	const iter = 150
	p2c := Point2Complex(lower, upper, m.Bounds())
	Iter(*m, func(x,y int) color.Color {
		factor := float64(Bounded(MandelbrotIteration, complex(0,0), p2c(image.Pt(x,y)), 2,iter))/iter
		//fmt.Print(int(factor*10))
		return color.RGBA{uint8(255*factor),uint8(255*factor),uint8(255*factor),255}
	})
	fo, err := os.Create("output.png")
	if err != nil {
		t.Error(err)
	}
	png.Encode(fo, m)
}