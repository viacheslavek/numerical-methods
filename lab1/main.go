package main

import (
	"fmt"

	"github.com/viacheslavek/numerical-methods/lab1/matrix"
	"github.com/viacheslavek/numerical-methods/lab1/running"
)

func main() {

	m := *matrix.NewMatrix[int](4, 4, [][]int{
		{4, 1, 0, 0},
		{1, 4, 1, 0},
		{0, 1, 4, 1},
		{0, 0, 1, 4},
	})
	coefficient := []int{5, 6, 6, 5}

	sle := *matrix.NewSLE[int, int](m, coefficient)

	ans := running.Solve[int, int](sle)

	fmt.Println("ANS:")
	for _, a := range ans {
		fmt.Print(a, " ")
	}

}
