package main

import "fmt"

func main() {
	var beratParsel int
	var kg, gram int
	var biayaKg, biayaGram, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	kg = beratParsel / 1000
	gram = beratParsel % 1000

	biayaKg = kg * 10000

	switch {
	case kg > 10:
		biayaGram = 0
	case gram >= 500:
		biayaGram = gram * 5
	case gram > 0:
		biayaGram = gram * 15
	default:
		biayaGram = 0
	}

	totalBiaya = biayaKg + biayaGram

	fmt.Println("Detail berat:", kg, "kg +", gram, "gram")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
