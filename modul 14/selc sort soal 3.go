package main

import "fmt"

type arrInt [1000000]int

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

func median(A arrInt, n int) int {

	if n%2 == 1 {
		return A[n/2]
	}

	return (A[(n/2)-1] + A[n/2]) / 2
}

func main() {

	var data arrInt
	var x int
	var n int

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {

			selectionSort(&data, n)
			fmt.Println(median(data, n))

		} else {

			data[n] = x
			n++

		}
	}
}