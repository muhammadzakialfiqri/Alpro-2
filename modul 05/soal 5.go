package main

import "fmt"

func ganjil(n, x int) {

	if x > n {
		return
	}

	if x%2 != 0 {
		fmt.Print(x, " ")
	}

	ganjil(n, x+1)

}

func main() {
	var n int

	fmt.Scan(&n)
	ganjil(n, 1)
}
