package main

import "fmt"

type arrInt [1000]int
type matrix [1000]arrInt

func selectionSort(A *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {

		idxMin = i

		for j = i + 1; j < n; j++ {

			if A[j] < A[idxMin] {
				idxMin = j
			}

		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {

	var data matrix
	var jumlah [1000]int
	var n int
	var i, j int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {

		fmt.Scan(&jumlah[i])

		for j = 0; j < jumlah[i]; j++ {
			fmt.Scan(&data[i][j])
		}
	}

	fmt.Println()

	for i = 0; i < n; i++ {

		selectionSort(&data[i], jumlah[i])

		for j = 0; j < jumlah[i]; j++ {

			if j == jumlah[i]-1 {
				fmt.Print(data[i][j])
			} else {
				fmt.Print(data[i][j], " ")
			}

		}

		fmt.Println()
	}
}