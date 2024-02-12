package running

import (
	"fmt"
	"log"

	"github.com/viacheslavek/numerical-methods/lab1/matrix"
)

type solver[M matrix.Additive, C matrix.Additive] struct {
	x     []float64
	y     []float64
	alpha []float64
	beta  []float64
	sle   matrix.SLE[M, C]
}

func Solve[M matrix.Additive, C matrix.Additive](sle matrix.SLE[M, C]) []float64 {

	fmt.Println(sle.Mat.String())

	fmt.Println(sle.Coefficient)

	s := solver[M, C]{
		x:     make([]float64, sle.Mat.GetRow()),
		y:     make([]float64, sle.Mat.GetRow()),
		alpha: make([]float64, sle.Mat.GetRow()),
		beta:  make([]float64, sle.Mat.GetRow()),
		sle:   sle,
	}

	s.straight()
	s.under()

	return s.x
}

func (s *solver[M, C]) straight() {

	a := make([]M, s.sle.Mat.GetRow())
	b := make([]M, s.sle.Mat.GetRow())
	c := make([]M, s.sle.Mat.GetRow())

	fillArray[M](a, b, c, s.sle.Mat)

	fmt.Println("a", a, "b", b, "c", c)

	s.y[0] = float64(b[0])
	s.alpha[0] = -float64(c[0]) / s.y[0]
	s.beta[0] = float64(s.sle.Coefficient[0]) / s.y[0]

	for i := 1; i < s.sle.Mat.GetRow(); i++ {
		s.y[i] = float64(b[i]) + float64(a[i])*s.alpha[i-1]

		s.alpha[i] = -float64(c[i]) / s.y[i]

		s.beta[i] = (float64(s.sle.Coefficient[i]) - float64(a[i])*s.beta[i-1]) / s.y[i]

		fmt.Println("i", i)
		fmt.Println("s.y[i]", s.y[i])
		fmt.Println("s.alpha[i]", s.alpha[i])
		fmt.Println("s.beta[i]", s.beta[i])
	}

	fmt.Println("y:", s.y)
	fmt.Println("alpha:", s.alpha)
	fmt.Println("beta:", s.beta)
	fmt.Println("x:", s.x)

}

func fillArray[T matrix.Additive](a, b, c []T, mat matrix.Matrix[T]) {
	for i := 0; i < mat.GetRow(); i++ {
		b[i] = mat.GetElem(i, i)
	}
	for i := 1; i < mat.GetRow(); i++ {
		a[i] = mat.GetElem(i, i-1)
	}
	for i := 0; i < mat.GetRow()-1; i++ {
		c[i] = mat.GetElem(i, i+1)
	}

	checkConditions[T](a, b, c, mat.GetRow())

}

func checkConditions[T matrix.Additive](a, b, c []T, N int) {
	if b[0] == 0 {
		log.Fatalf("b1 != 0")
	}

	if abs(c[0])/abs(b[0]) > 1 {
		log.Fatalf("|c1|/|b1| <= 1")
	}

	if abs(a[N-2])/abs(b[N-1]) > 1 {
		log.Fatalf("|a(n-1)|/|b(n)| <= 1")
	}

	for i := 1; i < N-1; i++ {
		if !(abs(b[i]) >= abs(a[i-1])+abs(c[i])) {
			log.Fatalf("|b(i)| >= |a(i-1)| + |c(i)|")
		}
	}
}

func abs[T matrix.Additive](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

func (s *solver[M, C]) under() {
	N := s.sle.Mat.GetRow() - 1
	s.x[N] = s.beta[N]

	for i := N - 1; i >= 0; i-- {
		s.x[i] = s.alpha[i]*s.x[i+1] + s.beta[i]
	}

	fmt.Println("final x", s.x)
}
