package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
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

func findEquationRootsOnSegmentNewtonMethod(currentPoint float64) float64 {
	iterations := 0

	for {
		if math.Abs(f(currentPoint)) < eps {
			fmt.Printf("Найдено методом Ньютона за %d итераций\n", iterations)
			return currentPoint
		}

		currentPoint -= f(currentPoint) / fx(currentPoint)

		iterations++
		if iterations > 10000 {
			fmt.Println("Слишком много итераций")
			return 0
		}
	}
}

func printEquationSolution() {

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
}

// INFO: Решение системы______________________________________________________

func fSystem(x *mat.VecDense) *mat.VecDense {
	result := mat.NewVecDense(2,
		[]float64{
			math.Cos(x.At(0, 0)-1) + x.At(1, 0) - 0.8,
			x.At(0, 0) - math.Cos(x.At(1, 0)) - 2,
		})
	return result
}

func jacobi(x *mat.VecDense) *mat.Dense {
	result := mat.NewDense(2, 2,
		[]float64{
			-math.Sin(x.At(0, 0) - 1), 1,
			1, math.Sin(x.At(1, 0)),
		})
	return result
}

func NewtonMethod(x0 *mat.VecDense) *mat.VecDense {
	x := x0

	for iterations := 0; iterations < 1000; iterations++ {
		dx := mat.NewVecDense(2, nil)
		err := dx.SolveVec(jacobi(x), fSystem(x))
		if err != nil {
			log.Fatalf("failed solve system %e", err)
			return nil
		}

		x.SubVec(x, dx)

		if mat.Norm(dx, 2) < eps {
			fmt.Printf("iterations: %d\n", iterations)
			return x
		}
	}

	log.Fatalf("too much iterations")
	return nil
}

func printSystemSolution() {
	initialPoint := []float64{2, 1}

	x0 := mat.NewVecDense(2, initialPoint)
	result := NewtonMethod(x0)

	fmt.Println("Решение: ", result.At(0, 0), result.At(1, 0))
}

func main() {
	fmt.Println("start program")

	printEquationSolution()

	fmt.Println("end program")
}
