package main

import "fmt"

func main() {
	var bunga string
	pita := ""
	jumlah := 0
	i := 1

	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" || bunga == "selesai" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
