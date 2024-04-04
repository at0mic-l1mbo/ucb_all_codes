package main

import (
	"fmt"
)

func transporMatriz(A [][]int) {
	fmt.Println("Matriz Original:")
	for _, coluna := range A {
		fmt.Println(coluna)
	}

	transposta := make([][]int, len(A[0]))
	for linha := range transposta {
		transposta[linha] = make([]int, len(A))
	}

	for linha, coluna := range A {
		for coluna, val := range coluna {
			transposta[coluna][linha] = val
		}
	}

	fmt.Println("\nMatriz Transposta:")
	for _, coluna := range transposta {
		fmt.Println(coluna)
	}
}

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	transporMatriz(A)
}
