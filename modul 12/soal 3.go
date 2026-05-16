package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {

	kr := 0
	kn := n - 1

	for kr <= kn {

		tengah := (kr + kn) / 2

		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kn = tengah - 1
		} else {
			kr = tengah + 1
		}
	}

	return -1
}

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
