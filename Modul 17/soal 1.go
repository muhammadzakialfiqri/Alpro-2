package main

import "fmt"

func main() {
	var bil, jumlah, rerata float64
	var banyak int

	fmt.Scan(&bil)

	for bil != 9999 {
		jumlah += bil
		banyak++
		fmt.Scan(&bil)
	}

	if banyak > 0 {
		rerata = jumlah / float64(banyak)
		fmt.Printf("%.2f\n", rerata)
	} else {
		fmt.Println("Tidak ada data")
	}
}