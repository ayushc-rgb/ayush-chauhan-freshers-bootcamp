package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Nrows    int
	Nrow     int
	Elements [][]int
}

func newMatrix(rows, cols int) Matrix {
	Elements := make([][]int, rows)
	for i := 0; i < rows; i++ {
		Elements[i] = make([]int, cols)
	}
	return Matrix{
		Nrows:    rows,
		Nrow:     cols,
		Elements: Elements,
	}
}
func (m *Matrix) getRows() int {
	return m.Nrows
}
func (m *Matrix) getCols() int {
	return m.Nrow
}
func (m *Matrix) setValue(x, y int, val int) error {
	m.Elements[x][y] = val
	if x < 0 || x >= m.Nrows || y < 0 || y >= m.Nrow {
		return fmt.Errorf("Index out of bounds")
	}
	m.Elements[x][y] = val
	return nil
}
func (m *Matrix) addMatrices(anotherMatrices Matrix) Matrix {
	rows := len(m.Elements)
	cols := len(m.Elements[0])
	result := newMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + anotherMatrices.Elements[i][j]
		}
	}
	return result
}
func (m *Matrix) printAsJSON() (string, error) {
	jsonData, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
func (m *Matrix) printMatrix() {
	for i := 0; i < m.Nrows; i++ {
		for j := 0; j < m.Nrow; j++ {
			fmt.Print(m.Elements[i][j], " ")

		}
		fmt.Println()
	}
}
func main() {
	matrixA := newMatrix(3, 3) // creating a 3x3 matrix
	matrixA.setValue(0, 0, 1)
	matrixA.setValue(0, 1, 2) // setting values in the matrix
	matrixA.setValue(0, 2, 3)
	matrixA.setValue(1, 0, 4)
	matrixA.setValue(1, 1, 5)
	matrixA.setValue(1, 2, 6)
	matrixA.setValue(2, 0, 7)
	matrixA.setValue(2, 1, 8)
	matrixA.setValue(2, 2, 9)

	// printing the matrix
	matrixA.printMatrix()
	fmt.Print("Number of Rows are:", matrixA.getRows())
	fmt.Print("Number of Columns are:", matrixA.getCols())

	// creating another matrix
	matrixB := newMatrix(3, 3)
	// Set some values in Matrix B.
	matrixB.setValue(0, 0, 9)
	matrixB.setValue(0, 1, 8)
	matrixB.setValue(0, 2, 7)
	matrixB.setValue(1, 0, 6)
	matrixB.setValue(1, 1, 5)
	matrixB.setValue(1, 2, 4)
	matrixB.setValue(2, 0, 3)
	matrixB.setValue(2, 1, 2)
	matrixB.setValue(2, 2, 1)

	matrixC := matrixA.addMatrices(matrixB) // so store the addition
	matrixC.printMatrix()
	// as json
	fmt.Println("As JSON")
	str, _ := matrixC.printAsJSON()
	fmt.Println(str)
}
