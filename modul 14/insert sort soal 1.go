package main

import "fmt"

type arrInt [1000]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {

		temp = A[i]
		j = i

		for j > 0 && temp < A[j-1] {

			A[j] = A[j-1]
			j--

		}

		A[j] = temp
	}
}

func main() {

	var data arrInt
	var x int
	var n int
	var i int
	var jarak int
	var tetap bool = true

	for {

		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data[n] = x
		n++

	}
		fmt.Println()

	insertionSort(&data, n)

	for i = 0; i < n; i++ {

		if i == n-1 {
			fmt.Print(data[i])
		} else {
			fmt.Print(data[i], " ")
		}

	}

	fmt.Println()

	jarak = data[1] - data[0]

	for i = 1; i < n-1; i++ {

		if data[i+1]-data[i] != jarak {
			tetap = false
		}

	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}