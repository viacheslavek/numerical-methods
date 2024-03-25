package main

import (
	"fmt"
	"math"
)

func direct(b, a, c, d []float64, size int) ([]float64, []float64) {
	alpha := make([]float64, 0, size)
	beta := make([]float64, 0, size)

	alpha = append(alpha, -c[0]/b[0])
	beta = append(beta, d[0]/b[0])

	for i := 1; i < size-1; i++ {
		alpha = append(alpha, -c[i]/(a[i-1]*alpha[i-1]+b[i]))
		beta = append(beta, (d[i]-a[i-1]*beta[i-1])/(a[i-1]*alpha[i-1]+b[i]))
	}

	beta = append(beta, (d[size-1]-a[size-2]*beta[size-2])/(a[size-2]*alpha[size-2]+b[size-1]))

	return alpha, beta
}

func reverse(alpha, beta []float64, size int) []float64 {
	x := make([]float64, 0, size)
	x = make([]float64, size)

	x[size-1] = beta[size-1]
	for i := size - 2; i >= 0; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}

	return x
}

// INFO: Для смены варианта изменить функцию, краевые условия и аналитическое решение
// Функция f(x)

func fx(x float64) float64 {
	return 2 * x
}

// Аналитическое решение
func solution(x float64) float64 {
	return x + math.Exp(x)*math.Sin(x) - math.Exp(x)*math.Cos(x) + 1
}

const (
	p          = -2.0
	q          = 2.0
	n          = 40
	outputStep = 4
)

var (
	y0    = 0.0
	dy0   = 1.0
	left  = solution(y0)
	right = solution(dy0)
	h     = 1.0 / float64(n)
)

func getA() []float64 {
	a := make([]float64, 0, n)
	for i := 1; i < n-1; i++ {
		a = append(a, 1-h/2*p)
	}
	return a
}

func getB() []float64 {
	b := make([]float64, 0, n)
	for i := 1; i < n; i++ {
		b = append(b, h*h*q-2)
	}
	return b
}

func getC() []float64 {
	c := make([]float64, 0, n)
	for i := 1; i < n-1; i++ {
		c = append(c, 1+h/2*p)
	}
	return c
}

func getD() []float64 {
	d := make([]float64, 0, n)
	d = append(d, h*h*fx(y0)-left*(dy0-h/2*p))

	for i := 2; i < n; i++ {
		d = append(d, h*h*fx(float64(i)*h))
	}

	d[len(d)-1] = h*h*fx(float64(len(d)-1)*h) - right*(dy0+h/2*p)

	return d
}

func getX() []float64 {
	x := make([]float64, 0, n)
	for i := 0; i < n+1; i++ {
		x = append(x, float64(i)*h)
	}
	return x
}

func getY(a, b, c, d []float64) []float64 {
	alpha, beta := direct(b, a, c, d, n-1)

	y := make([]float64, 0, n)
	y = append(y, left)
	y = append(y, reverse(alpha, beta, n-1)...)
	y = append(y, right)

	return y
}

func diffSolution() ([]float64, []float64) {
	return getX(), getY(getA(), getB(), getC(), getD())
}

func main() {
	fmt.Printf("Input:\n")
	fmt.Printf("y'' + (%0.1f)*y' + (%0.1f)*y = 2*x\n", p, q)
	fmt.Printf("y(0) = %f\ny(1) = %f\n", left, right)
	fmt.Println()

	fmt.Printf("Solution:\n")
	x, y := diffSolution()

	for i := 0; i < len(x); i += outputStep {
		fmt.Printf("x=%.2f, y=%.6f, y*=%.6f  |y-y*|=%.6f\n",
			float64(i)*h, solution(x[i]), y[i], math.Abs(y[i]-solution(x[i])))
	}

}
