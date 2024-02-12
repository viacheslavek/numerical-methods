package matrix

import (
	"fmt"
	"log"
)

type Additive interface {
	int | uint | ~uint64 | ~float64
}

type Matrix[T Additive] struct {
	row, col int
	buffer   [][]T
}

func NewMatrix[T Additive](rows, cols int, buffer [][]T) *Matrix[T] {

	m := &Matrix[T]{
		row: rows,
		col: cols,
	}

	m.buffer = make([][]T, rows)

	for i := 0; i < rows; i++ {
		m.buffer[i] = make([]T, cols)
		for j := 0; j < cols; j++ {
			m.buffer[i][j] = buffer[i][j]
		}
	}

	return m
}

func (m *Matrix[T]) GetRow() int {
	return m.row
}

func (m *Matrix[T]) GetCol() int {
	return m.col
}

func (m *Matrix[T]) GetBuffer() [][]T {
	return m.buffer
}

func (m *Matrix[T]) GetElem(i, j int) T {
	if i > m.GetRow() || j > m.GetCol() {
		fmt.Println("out of range")
		var temp T
		return temp
	}
	return m.buffer[i][j]
}

func (m *Matrix[T]) String() string {
	ans := "[\n"

	for i := 0; i < m.row; i++ {
		ans += fmt.Sprintln(m.buffer[i])
	}

	ans += "]"

	return ans
}

// SLE system of linear equations
type SLE[M, C Additive] struct {
	Mat         Matrix[M]
	Coefficient []C
}

func NewSLE[M, C Additive](m Matrix[M], coefficient []C) *SLE[M, C] {

	if m.row != len(coefficient) {
		log.Fatal("rows don`t equal len of coefficient")
	}

	return &SLE[M, C]{
		Mat:         m,
		Coefficient: coefficient,
	}
}
