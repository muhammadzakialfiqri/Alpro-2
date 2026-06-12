package main

import "fmt"

func main() {
	var n int
	var jumlah float64

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			jumlah = jumlah + 1.0/float64(2*i-1)
		} else {
			jumlah = jumlah - 1.0/float64(2*i-1)
		}
	}

	fmt.Printf("Hasil PI: %.7f\n", 4*jumlah)
}