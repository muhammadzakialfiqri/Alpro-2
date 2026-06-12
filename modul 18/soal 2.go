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
		return Domino{}
	}

	kartu := d.Kartu[d.Jumlah-1]
	d.Jumlah--

	return kartu
}

func galiKartu(d *Dominoes, target Domino) {

	fmt.Printf(
		"\nMencari kartu yang memiliki suit %d atau %d\n",
		target.Sisi1,
		target.Sisi2,
	)

	for d.Jumlah > 0 {

		kartu := ambilKartu(d)

		fmt.Printf(
			"Mengambil (%d,%d)\n",
			kartu.Sisi1,
			kartu.Sisi2,
		)

		if kartu.Sisi1 == target.Sisi1 ||
			kartu.Sisi1 == target.Sisi2 ||
			kartu.Sisi2 == target.Sisi1 ||
			kartu.Sisi2 == target.Sisi2 {

			fmt.Println("Kartu cocok ditemukan!")
			fmt.Printf(
				"Kartu = (%d,%d)\n",
				kartu.Sisi1,
				kartu.Sisi2,
			)

			return
		}
	}

	fmt.Println("Tidak ada kartu yang cocok.")
}

func sepasangKartu(d1 Domino, d2 Domino) bool {

	return d1.Nilai+d2.Nilai == 12
}

func main() {

	dominoes := buatDominoes()

	kocokKartu(&dominoes)

	var a, b int

	fmt.Println("")

	fmt.Print("Masukkan sisi pertama kartu target : ")
	fmt.Scan(&a)

	fmt.Print("Masukkan sisi kedua kartu target : ")
	fmt.Scan(&b)

	target := Domino{
		Sisi1: a,
		Sisi2: b,
		Nilai: a + b,
		Balak: a == b,
	}

	galiKartu(&dominoes, target)

	fmt.Println("")

	k1 := ambilKartu(&dominoes)
	k2 := ambilKartu(&dominoes)

	fmt.Printf(
		"Kartu 1 : (%d,%d) Nilai=%d\n",
		k1.Sisi1,
		k1.Sisi2,
		k1.Nilai,
	)

	fmt.Printf(
		"Kartu 2 : (%d,%d) Nilai=%d\n",
		k2.Sisi1,
		k2.Sisi2,
		k2.Nilai,
	)

	fmt.Println(
		"Apakah total nilainya 12?",
		sepasangKartu(k1, k2),
	)
}