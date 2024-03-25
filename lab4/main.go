package main

import (
	"fmt"
	"log"
	"math"
)

func getSum(nums []float64) float64 {
	summer := 0.0
	for _, num := range nums {
		summer += num
	}
	return summer
}

func getInverseArr(nums []float64) []float64 {
	inverse := make([]float64, len(nums))
	for i, num := range nums {
		inverse[i] = 1 / num
	}
	return inverse
}

func getSquaresArr(nums []float64) []float64 {
	squares := make([]float64, len(nums))
	for i, num := range nums {
		squares[i] = num * num
	}
	return squares
}

func getDivideFirstArrBySecond(first, second []float64) []float64 {
	if len(first) != len(second) {
		log.Fatalf("len first != len second")
	}
	dividers := make([]float64, len(first))

	for i := 0; i < len(first); i++ {
		dividers[i] = first[i] / second[i]
	}
	return dividers
}

func getCoefficient(x, y []float64,
	f func(ySum, xInverseSum, xInverseSquaresSum, yDivideByXSum, N float64) float64) float64 {

	if len(x) != len(y) {
		log.Fatalf("len x != len y")
	}

	N := float64(len(x))
	xInverse := getInverseArr(x)
	xInverseSquares := getSquaresArr(xInverse)
	yDivideByX := getDivideFirstArrBySecond(y, x)

	ySum := getSum(y)
	xInverseSum := getSum(xInverse)
	xInverseSquaresSum := getSum(xInverseSquares)
	yDivideByXSum := getSum(yDivideByX)

	return f(ySum, xInverseSum, xInverseSquaresSum, yDivideByXSum, N)
}

func getACoefficient(x, y []float64) float64 {
	return getCoefficient(
		x, y,
		func(ySum, xInverseSum, xInverseSquaresSum, yDivideByXSum, N float64) float64 {
			return (ySum*xInverseSquaresSum - xInverseSum*yDivideByXSum) /
				(N*xInverseSquaresSum - xInverseSum*xInverseSum)
		})
}

func getBCoefficient(x, y []float64) float64 {
	return getCoefficient(
		x, y,
		func(ySum, xInverseSum, xInverseSquaresSum, yDivideByXSum, N float64) float64 {
			return (N*yDivideByXSum - xInverseSum*ySum) /
				(N*xInverseSquaresSum - xInverseSum*xInverseSum)
		})
}

func getSCU(a, b float64, x, y []float64) float64 {
	scuElements := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		scuElements[i] = math.Abs(a + b/x[i] - y[i])
	}

	scuSquaresElements := getSquaresArr(scuElements)

	return getSum(scuSquaresElements)
}

func getSCO(scu float64, N float64) float64 {
	return scu / math.Sqrt(N)
}

func main() {

	x := []float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}
	y := []float64{2.61, 1.62, 1.17, 0.75, 0.30, 0.75, 1.03, 0.87, 0.57}

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	a := getACoefficient(x, y)
	b := getBCoefficient(x, y)
	fmt.Printf("Коэффициенты a: %f и b: %f\n", a, b)

	fmt.Printf("Итоговая функция: y = %f + %f / x\n", a, b)

	scu := getSCU(a, b, x, y)
	fmt.Printf("СКУ: %f\n", scu)

	sco := getSCO(scu, float64(len(x)))
	fmt.Printf("Средняя ошибка апроксимации: %f\n", sco)

}
