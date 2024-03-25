package main

import (
	"fmt"
	"math"
)

const (
	epsilon                   = 0.001
	lowBorder                 = 1.0 / math.E
	upBorder                  = math.E
	accuracyStepForRichardson = 4
)

// INFO: для смены варианта нужно изменить эту функцию

func f(x float64) float64 {
	return math.Pow(math.Log(x), 2) / x
}

func rectangle(a float64, b float64, n int) float64 {
	h := (b - a) / float64(n)
	var s float64
	for i := 1; i <= n; i++ {
		s += f(a + (float64(i)-0.5)*h)
	}
	return h * s
}

func trapezoid(a float64, b float64, n int) float64 {
	h := (b - a) / float64(n)
	var s float64
	for i := 1; i < n; i++ {
		s += f(a + float64(i)*h)
	}
	return h * ((f(a)+f(b))/2 + s)
}

func simpson(a float64, b float64, n int) float64 {
	h := (b - a) / float64(n)

	var s1, s2, s3 float64
	for i := 1; i <= n; i++ {
		s1 += f(a + float64(i)*h)
	}
	for i := 1; i <= n; i++ {
		s2 += f(a + (float64(i)-0.5)*h)
	}
	for i := 1; i <= n; i++ {
		s3 += f(a + (float64(i)-1)*h)
	}

	s := s1 + 4*s2 + s3

	return h / 6 * s
}

func getIntegralValue(
	calculate func(float64, float64, int) float64,
	a float64, b float64) (int, float64, float64) {
	n := 1
	richardson := epsilon * 1000
	result := 0.0
	i := 0

	for math.Abs(richardson) >= epsilon {
		n *= 2
		prevResult := result
		result = calculate(a, b, n)
		richardson = (result - prevResult) / (math.Pow(2, accuracyStepForRichardson) - 1)
		i++
	}

	return n, result, richardson
}

func printAll(
	nRec, nTra, nSim int,
	resRec, resTra, resSim float64,
	richRec, richTra, richSim float64) {

	fmt.Printf("\n")
	fmt.Printf("\tCentral rectangles  \tTrapezoids method   \tSimpson's\n")

	fmt.Printf("n:\t")
	fmt.Printf("%d\t\t", nRec)
	fmt.Printf("\t%d\t\t", nTra)
	fmt.Printf("\t%d\t\t", nSim)
	fmt.Println()
	fmt.Printf("I*:\t")
	fmt.Printf("%.20f\t", resRec)
	fmt.Printf("%.20f\t", resTra)
	fmt.Printf("%.20f\t", resSim)
	fmt.Println()
	fmt.Printf("R:\t")
	fmt.Printf("%.20f\t", richRec)
	fmt.Printf("%.20f\t", richTra)
	fmt.Printf("%.20f\t", richSim)
	fmt.Println()
	fmt.Printf("I* + R: ")
	fmt.Printf("%.20f\t", resRec+richRec)
	fmt.Printf("%.20f\t", resTra+richTra)
	fmt.Printf("%.20f", resSim+richSim)
	fmt.Println()

}

func main() {

	fmt.Println("epsilon =", epsilon, "| analytic ans = ", 2.0/3.0)

	nRec, resRec, richRec := getIntegralValue(rectangle, lowBorder, upBorder)

	nTra, resTra, richTra := getIntegralValue(trapezoid, lowBorder, upBorder)

	nSim, resSim, richSim := getIntegralValue(simpson, lowBorder, upBorder)

	printAll(nRec, nTra, nSim, resRec, resSim, resTra, richRec, richTra, richSim)

}
