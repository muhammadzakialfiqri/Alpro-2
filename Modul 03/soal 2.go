package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("masukan nilai a b dan c : ")
	fmt.Scan(&a, &b, &c)

	fmt.Printf("fogog (%d) : %d\n", a, f(g(h(a))))
	fmt.Printf("gohof (%d) : %d\n", b, g(h(f(b))))
	fmt.Printf("hofog (%d) : %d\n", c, h(f(g(c))))

}

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}
