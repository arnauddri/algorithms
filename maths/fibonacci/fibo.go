// Package main provides ...
package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := 10000
	simul := 10000
	memo := make(map[int][4]int)
	recursivePts := make(plotter.XYs, n)
	iterativePts := make(plotter.XYs, n)

	for i := 11; i < n; i += 10 {
		tm := 0.0
		ti := 0.0
		elapsed := time.Duration(0)
		for j := 0; j < simul; j++ {
			t0 := time.Now()
			getNumber(rand.Intn(i-10)+10, memo)
			elapsed = time.Since(t0)
			tm += elapsed.Seconds()

			t1 := time.Now()
			getFiboIterative(rand.Intn(i-10) + 10)
			elapsed = time.Since(t1)
			ti += elapsed.Seconds()
		}
		recursivePts[i].X = float64(i)
		recursivePts[i].Y = tm / float64(simul)
		iterativePts[i].X = float64(i)
		iterativePts[i].Y = ti / float64(simul)
	}
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Calculation of Fibonacci numbers"
	p.X.Label.Text = "Approximate number n"
	p.Y.Label.Text = "Running Time"

	err = plotutil.AddLines(p,
		"Recursive", recursivePts,
		"Iterative", iterativePts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(14, 10, "fibo.png"); err != nil {
		panic(err)
	}
}

func getNumber(n int, memo map[int][4]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	power := string(strconv.FormatInt(int64(n-1), 2))
	powers := make([]int, 0)

	for i, r := range power {
		if string(r) == "1" {
			powers = append([]int{int(math.Pow(2, float64(len(power)-i-1)))}, powers...)
		}
	}

	Q := [4]int{1, 1, 1, 0}
	matrices := make([][4]int, 0)

	for _, p := range powers {
		matrices = append(matrices, raiseToPower(Q, p, memo))
	}

	for len(matrices) > 1 {
		M1 := matrices[0]
		M2 := matrices[1]
		R := multiplyMatrix(M1, M2)
		matrices = matrices[2:len(matrices)]
		matrices = append(matrices, R)
	}

	return matrices[0][0]
}

func multiplyMatrix(A [4]int, B [4]int) [4]int {
	var result [4]int

	result[0] = A[0]*B[0] + A[1]*B[2]
	result[1] = A[0]*B[1] + A[1]*B[3]
	result[2] = A[2]*B[0] + A[3]*B[2]
	result[3] = A[2]*B[1] + A[3]*B[3]

	return result
}

func raiseToPower(A [4]int, p int, memo map[int][4]int) [4]int {
	var result [4]int

	if p == 1 {
		return A
	}

	if cache, ok := memo[p]; ok {
		return cache
	}

	result = raiseToPower(A, p/2, memo)
	R := multiplyMatrix(result, result)

	memo[p] = R

	return R
}

func getFiboIterative(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	for i, j := 0, 1; j < n-1; i, j = i+j, i {
		a = j
	}

	return a
}
