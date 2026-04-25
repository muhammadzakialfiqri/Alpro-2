package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	fmt.Println()

	var i int = 0
	var jumlahWadah int = 0
	var beratWadah float64 = 0

	for i < x {
		var total float64 = 0

		for j := 0; j < y && i < x; j++ {
			total += ikan[i]
			i++
		}

		fmt.Print(total, " ")

		beratWadah = beratWadah + total
		jumlahWadah = jumlahWadah + 1
	}

	fmt.Println()

	rataRata := beratWadah / float64(jumlahWadah)
	fmt.Printf("%.2f",rataRata)
}