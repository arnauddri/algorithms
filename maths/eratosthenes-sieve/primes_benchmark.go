package primes

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"testing"
)

func benchmarkGetAllPrimesTo(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAllPrimesTo(n)
	}
}

func BenchmarkGetAllPrimesTo10(b *testing.B)         { benchmarkGetAllPrimesTo(10, b) }
func BenchmarkGetAllPrimesTo100(b *testing.B)        { benchmarkGetAllPrimesTo(100, b) }
func BenchmarkGetAllPrimesTo1000(b *testing.B)       { benchmarkGetAllPrimesTo(1000, b) }
func BenchmarkGetAllPrimesTo10000(b *testing.B)      { benchmarkGetAllPrimesTo(10000, b) }
func BenchmarkGetAllPrimesTo100000(b *testing.B)     { benchmarkGetAllPrimesTo(100000, b) }
func BenchmarkGetAllPrimesTo1000000(b *testing.B)    { benchmarkGetAllPrimesTo(1000000, b) }
func BenchmarkGetAllPrimesTo10000000(b *testing.B)   { benchmarkGetAllPrimesTo(10000000, b) }
func BenchmarkGetAllPrimesTo100000000(b *testing.B)  { benchmarkGetAllPrimesTo(100000000, b) }
func BenchmarkGetAllPrimesTo1000000000(b *testing.B) { benchmarkGetAllPrimesTo(1000000000, b) }

func TestBenchmark(t *testing.T) {
	points := make(plotter.XYs, 9)

	points[0].X = 10
	points[1].X = 100
	points[2].X = 1000
	points[3].X = 10000
	points[4].X = 100000
	points[5].X = 1000000
	points[6].X = 10000000
	points[7].X = 100000000
	points[8].X = 1000000000
	points[0].Y = 496.0
	points[1].Y = 1196
	points[2].Y = 7724
	points[3].Y = 58489
	points[4].Y = 414589
	points[5].Y = 5590372
	points[6].Y = 96922714
	points[7].Y = 1431869649
	points[8].Y = 24368113980

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Calculation of the primes between 0 and n"
	p.X.Label.Text = "Primes Range"
	p.Y.Label.Text = "Running Time (ns/op)"

	err = plotutil.AddLines(p,
		"Sieve", points)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(14, 10, "benchmark.png"); err != nil {
		panic(err)
	}
}
