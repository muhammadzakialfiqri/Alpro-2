package main

import "fmt"

func barisanNaikTurun(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")

	barisanNaikTurun(n - 1)

	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	barisanNaikTurun(n)
}