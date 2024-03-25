package main

import (
	"fmt"
	"math"
)

//const (
//	N     = 10
//	Left  = 0.0
//	Right = 1.0
//)

// Вариант 17

const (
	Left  = 1.0 / math.E
	Right = math.E
	N     = 32
)

func testFunc(x float64) float64 {
	return x*x*x + 2*x*x - 3*x + 4
}

func varFunc(x float64) float64 {
	return math.Pow(math.Log(x), 2) / x
}

func getXY(h float64, f func(x float64) float64) ([]float64, []float64) {
	xArr := make([]float64, N+1)
	for i := 0; i <= N; i++ {
		xArr[i] = Left + float64(i)*h
	}

	yArr := make([]float64, N+1)
	for i := 0; i <= N; i++ {
		yArr[i] = f(xArr[i])
	}

	return xArr, yArr
}

func printTableForFunc(x, y []float64) {
	fmt.Println("table for f(x):")
	for i := 0; i <= N; i++ {
		fmt.Printf("%.1f; %.16f\n", x[i], y[i])
	}
}

func getMatrixCoffs(y []float64, h float64) ([]float64, []float64, []float64, []float64) {

	a := make([]float64, 0)
	for i := 1; i < N-1; i++ {
		a = append(a, 1)
	}

	b := make([]float64, 0)
	for i := 1; i < N; i++ {
		b = append(b, 4)
	}

	c := make([]float64, 0)
	for i := 1; i < N-1; i++ {
		c = append(c, 1)
	}

	d := make([]float64, 0)
	for i := 1; i < N; i++ {
		d = append(d, 3*(y[i+1]-2*y[i]+y[i-1])/(h*h))
	}

	return a, b, c, d
}

func getA(y []float64) []float64 {
	A := make([]float64, 0)
	for i := 1; i <= N; i++ {
		A = append(A, y[i-1])
	}
	return A
}

func getB(y, C []float64, h float64) []float64 {
	B := make([]float64, 0)
	for i := 1; i <= N; i++ {
		B = append(B, (y[i]-y[i-1])/h-(h/3)*(C[i]+2*C[i-1]))
	}
	return B
}

func getAlphaBeta(a, b, c, d []float64, size int) (alpha, beta []float64) {
	alpha = append(alpha, -c[0]/b[0])
	beta = append(beta, d[0]/b[0])

	var y float64

	for i := 1; i < size-1; i++ {
		y = a[i-1]*alpha[i-1] + b[i]
		alpha = append(alpha, -c[i]/y)
		beta = append(beta, (d[i]-a[i-1]*beta[i-1])/y)
	}

	y = a[size-2]*alpha[size-2] + b[size-1]

	beta = append(beta, (d[size-1]-a[size-2]*beta[size-2])/y)

	return alpha, beta
}

func getC(y []float64, h float64) []float64 {

	a, b, c, d := getMatrixCoffs(y, h)
	alpha, beta := getAlphaBeta(a, b, c, d, N-1)

	C := make([]float64, N-1)
	C[N-2] = beta[N-2]

	for i := N - 3; i >= 0; i-- {
		C[i] = alpha[i]*C[i+1] + beta[i]
	}

	C = append(C, 0)
	C = append([]float64{0}, C...)

	return C
}

func getD(C []float64, h float64) []float64 {
	D := make([]float64, 0)
	for i := 1; i <= N; i++ {
		D = append(D, (C[i]-C[i-1])/(3*h))
	}
	return D
}

func PrintInterpolationNodes(A, B, C, D []float64, h float64, xArr []float64, f func(x float64) float64) {
	fmt.Println("interpolation nodes:")
	for i := 0; i < N; i++ {
		varX := Left + float64(i)*h
		varY := f(varX)
		s := A[i] + B[i]*(varX-xArr[i]) + C[i]*math.Pow(varX-xArr[i], 2) + D[i]*math.Pow(varX-xArr[i], 3)
		fmt.Printf("i: %d, x: %.2f, f(x): %.16f, y: %.16f, |f(x)-y2|: %.16f\n",
			i+1, varX, s, varY, math.Abs(s-varY))
	}
}

func PrintMiddleInterpolationNodes(A, B, C, D []float64, h float64, xArr []float64, f func(x float64) float64) {
	fmt.Println("middles interpolation nodes:")
	for i := 0; i < N; i++ {
		varX := Left + (float64(i+1)-0.5)*h
		varY := f(varX)
		s := A[i] + B[i]*(varX-xArr[i]) + C[i]*math.Pow(varX-xArr[i], 2) + D[i]*math.Pow(varX-xArr[i], 3)
		fmt.Printf("i: %d, x: %.2f, f(x): %.16f, y: %.16f, |f(x)-y|: %.16f\n",
			i+1, varX, s, varY, math.Abs(s-varY))
	}
}

func main() {

	// INFO: Для смены варианта надо поменять константы Left, Right, N и функцию вычисления

	var h = (Right - Left) / N

	xArr, yArr := getXY(h, varFunc)

	printTableForFunc(xArr, yArr)

	A := getA(yArr)
	C := getC(yArr, h)
	B := getB(yArr, C, h)
	D := getD(C, h)

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("D:", D)

	PrintInterpolationNodes(A, B, C, D, h, xArr, varFunc)
	PrintMiddleInterpolationNodes(A, B, C, D, h, xArr, varFunc)

}
