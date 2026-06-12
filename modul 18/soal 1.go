package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	Kartu  [28]Domino
	Jumlah int
}

func buatDominoes() Dominoes {

	var d Dominoes
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {

			d.Kartu[idx] = Domino{
				Sisi1: i,
				Sisi2: j,
				Nilai: i + j,
				Balak: i == j,
			}

			idx++
		}
	}

	d.Jumlah = idx

	return d
}

func kocokKartu(d *Dominoes) {

	rand.Seed(time.Now().UnixNano())

	for i := d.Jumlah - 1; i > 0; i-- {

		j := rand.Intn(i + 1)

		d.Kartu[i], d.Kartu[j] =
			d.Kartu[j], d.Kartu[i]
	}
}


func ambilKartu(d *Dominoes) Domino {

	if d.Jumlah == 0 {
		fmt.Println("Kartu sudah habis!")
		return Domino{}
	}

	kartu := d.Kartu[d.Jumlah-1]
	d.Jumlah--

	return kartu
}

func gambarKartu(d Domino, suit int) int {

	if suit == 1 {
		return d.Sisi1
	}

	return d.Sisi2
}

func nilaiKartu(d Domino) int {
	return d.Nilai
}

func main() {

	dominoes := buatDominoes()

	fmt.Println("Jumlah kartu awal =", dominoes.Jumlah)

	kocokKartu(&dominoes)

	var n int

	fmt.Print("Berapa kartu yang ingin diambil? ")
	fmt.Scan(&n)

	if n > dominoes.Jumlah {
		n = dominoes.Jumlah
	}

	fmt.Println("\nHASIL PENGAMBILAN KARTU")

	for i := 1; i <= n; i++ {

		kartu := ambilKartu(&dominoes)

		fmt.Println("--------------------------------")
		fmt.Println("Kartu ke-", i)
		fmt.Printf("Domino      : (%d,%d)\n",
			kartu.Sisi1,
			kartu.Sisi2)

		fmt.Println("Suit 1      :", gambarKartu(kartu, 1))
		fmt.Println("Suit 2      :", gambarKartu(kartu, 2))
		fmt.Println("Nilai Kartu :", nilaiKartu(kartu))
		fmt.Println("Balak       :", kartu.Balak)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Sisa kartu :", dominoes.Jumlah)
}