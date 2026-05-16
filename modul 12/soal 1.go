package main

import "fmt"

func main() {

	var suara [21]int
	var x, Suaramasuk, Suarasah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}
		Suaramasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			Suarasah++
		}
	}

	fmt.Println("Suara masuk:", Suaramasuk)
	fmt.Println("Suara sah:", Suarasah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}
