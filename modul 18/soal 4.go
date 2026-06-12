package main

import (
	"fmt"
	"strings"
)

var pita string
var posisi int
var currentChar byte

func start(teks string) {
	pita = teks
	posisi = 0

	if len(pita) > 0 {
		currentChar = pita[posisi]
	}
}

func maju() {

	posisi++

	if posisi < len(pita) {
		currentChar = pita[posisi]
	}
}

func eop() bool {

	if posisi >= len(pita) {
		return true
	}

	return currentChar == '.'
}

func cc() byte {
	return currentChar
}

func main() {

	var input string

	fmt.Print("Masukkan kalimat (akhiri dengan titik): ")
	fmt.Scanln(&input)

	input = strings.ToUpper(input)

	start(input)

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	fmt.Println("\nKarakter yang terbaca:")

	for !eop() {

		fmt.Printf("%c ", cc())

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if posisi+1 < len(pita) {

			if pita[posisi] == 'L' &&
				pita[posisi+1] == 'E' {

				jumlahLE++
			}
		}

		maju()
	}

	fmt.Println()

	frekuensiA := 0.0

	if jumlahKarakter > 0 {

		frekuensiA =
			float64(jumlahA) /
				float64(jumlahKarakter) * 100
	}

	fmt.Println("\n===== HASIL =====")
	fmt.Println("Jumlah Karakter :", jumlahKarakter)
	fmt.Println("Jumlah Huruf A  :", jumlahA)
	fmt.Printf("Frekuensi A     : %.2f%%\n", frekuensiA)
	fmt.Println("Jumlah LE       :", jumlahLE)
}