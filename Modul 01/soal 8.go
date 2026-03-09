package main

import "fmt"

func main() {
	var k1, k2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&k1, &k2)

		if k1 < 0 || k2 < 0 || k1+k2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if k1-k2 >= 9 || k2-k1 >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
