package main

import "fmt"

func main() {

	var r, volume, luas float64
	const phi = 3.1415926535

	fmt.Print("Masukan jari-jari : ")
	fmt.Scan(&r)

	volume = 4.0 / 3.0 * phi * r * r * r
	luas = 4 * phi * r * r

	fmt.Printf("bola dengan jari-jari  %.f memiliki volumebola %.4f dan luas kulit %.4f ", r, volume, luas)
}
