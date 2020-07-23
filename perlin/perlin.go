// package perlin

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {

// 	generateLines()
// }

// func generateLines() {
// 	rand.Seed(time.Now().UnixNano())

// 	ret := "X"
// 	for x := 0; x < 10; x++ {
// 		no := rand.Intn(10)
// 		// no := 3
// 		lines := ""
// 		for y := 0; y < no; y++ {
// 			lines += "\n"
// 		}
// 		ret += lines
// 		for z := 0; z < x; z++ {
// 			ret += "\t"
// 		}
// 		ret += "X"
// 	}
// 	ret += "\n"
// 	fmt.Printf(ret)
// }

// package main

// import (
// 	"image/color"
// 	"log"
// 	"math/rand"
// 	"time"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	// randomPoints returns some random x, y points
// 	// with some interesting kind of trend.
// 	randomPoints := func(n int) plotter.XYs {
// 		pts := make(plotter.XYs, n)
// 		// for i := range pts {
// 		// 	if i == 0 {
// 		// 		pts[i].X = rnd.Float64()
// 		// 	} else {
// 		// 		pts[i].X = pts[i-1].X + rnd.Float64()
// 		// 	}
// 		// 	pts[i].Y = pts[i].X + 10*rnd.Float64()
// 		// }
// 		for i := 0; i < n; i++ {
// 			pts[i].X = float64(i)
// 			pts[i].Y = float64(rnd.Float32())
// 		}
// 		return pts
// 	}

// 	n := 100
// 	scatterData := randomPoints(n)
// 	// scatterData := plotter.XYs{
// 	// 	{
// 	// 		X: 10,
// 	// 		Y: 5,
// 	// 	},
// 	// }
// 	lineData := randomPoints(n)
// 	// linePointsData := randomPoints(n)

// 	p, err := plot.New()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	p.Title.Text = "Points Example"
// 	p.X.Label.Text = "X"
// 	p.Y.Label.Text = "Y"
// 	p.Add(plotter.NewGrid())

// 	s, err := plotter.NewScatter(scatterData)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
// 	s.GlyphStyle.Radius = vg.Points(3)

// 	l, err := plotter.NewLine(lineData)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	l.LineStyle.Width = vg.Points(1)
// 	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
// 	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

// 	// lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }
// 	// lpLine.Color = color.RGBA{G: 255, A: 255}
// 	// lpPoints.Shape = draw.CircleGlyph{}
// 	// lpPoints.Color = color.RGBA{R: 255, A: 255}

// 	p.Add(s)
// 	// p.X.Max = float64(n)
// 	// p.Y.Max = 10

// 	p.Legend.Add("scatter", s)
// 	// p.Legend.Add("line", l)
// 	// p.Legend.Add("line points", lpLine, lpPoints)

// 	err = p.Save(600, 200, "testdata/scatter.png")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
