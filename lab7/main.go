package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const eps = 0.001

func f(x float64) float64 {
	return x*x*x + x*x - 7*x + 4
}

func rootsF() []float64 {
	return []float64{-3.4027, 0.68397, 1.7187}
}

func printF() string {
	return "x*x*x + x*x - 7*x + 4"
}

func fx(x float64) float64 {
	return 3*x*x + 2*x - 7
}

func printFx() string {
	return "3*x*x + 2*x - 7"
}

func fxx(x float64) float64 {
	return 6*x + 2
}

func printFxx() string {
	return "6*x + 2"
}

func plotGraphic(g func(x float64) float64, gPrint func() string, left, right, down, up float64) {
	p := plot.New()

	p.Title.Text = "Graph of " + gPrint()
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	countPoints := 1000

	points := make(plotter.XYs, countPoints)
	for i := range points {
		x := left + ((right-left)/float64(countPoints))*float64(i)
		y := g(x)
		points[i].X = x
		points[i].Y = y
	}

	line, scatter, err := plotter.NewLinePoints(points)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.Add(line, scatter)
	p.X.Min = left
	p.X.Max = right
	p.Y.Min = down
	p.Y.Max = up

	if errS := p.Save(4*vg.Inch, 4*vg.Inch, "\""+gPrint()+"\""+"_graph.png"); err != nil {
		fmt.Println("failed to save png:", errS)
		return
	}
}

func findEquationRootsOnSegmentBisectionMethod(left, right float64) float64 {
	counter := 0

	for math.Abs(right-left) > eps {
		mid := (left + right) / 2
		if f(mid)*f(left) < 0 {
			right = mid
		} else {
			left = mid
		}
		counter++
	}

	fmt.Printf("Найдено методом деления пополам за %d делений\n", counter)

	return (right + left) / 2
}

func findEquationRootsOnSegmentNewtonMethod(initialPoint float64) float64 {

	iterations := 0

	for {
		fxVal := fx(initialPoint)
		if math.Abs(fxVal) < eps {
			break
		}
		fxxVal := fxx(initialPoint)
		if math.Abs(fxxVal) < eps {
			break
		}
		initialPoint = initialPoint - fxVal/fxxVal

		iterations++

		if iterations > 10000 {
			fmt.Println("Слишком много итераций")
			return 0
		}
	}

	fmt.Printf("Найдено методом Ньютона за %d итераций\n", iterations)

	return initialPoint
}

func main() {
	fmt.Println("start program")

	plotGraphic(f, printF, -5, 5, -30, 30)

	plotGraphic(fx, printFx, -3, 3, -8, 2)

	plotGraphic(fxx, printFxx, -1, 1, -1, 4)

	firstRootBis := findEquationRootsOnSegmentBisectionMethod(-4, 0)
	secondRootBis := findEquationRootsOnSegmentBisectionMethod(0, 1)
	thirdRootBis := findEquationRootsOnSegmentBisectionMethod(1, 4)

	fmt.Println("firstRootBis:", firstRootBis)
	fmt.Println("secondRootBis:", secondRootBis)
	fmt.Println("thirdRootBis:", thirdRootBis)

	firstRootNew := findEquationRootsOnSegmentNewtonMethod(-3)
	secondRootNew := findEquationRootsOnSegmentNewtonMethod(0.5)
	thirdRootNew := findEquationRootsOnSegmentNewtonMethod(2)

	fmt.Println("firstRootNew:", firstRootNew)
	fmt.Println("secondRootNew:", secondRootNew)
	fmt.Println("thirdRootNew:", thirdRootNew)

	fmt.Println("end program")
}