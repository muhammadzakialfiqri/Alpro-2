package main

import "fmt"

func cetakderet(n int) {
	for {
		fmt.Print(n, " ")

		if n == 1 {
			break
		} else if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	cetakderet(a)
}
