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

func bisaMain(k Domino, ujung int) bool {

	return k.Sisi1 == ujung ||
		k.Sisi2 == ujung
}

func hapusKartu(
	kartu []Domino,
	idx int,
) []Domino {

	return append(
		kartu[:idx],
		kartu[idx+1:]...,
	)
}

func tampilkanKartu(kartu []Domino) {

	fmt.Println("\nKartu Pemain")

	for i := 0; i < len(kartu); i++ {

		fmt.Printf(
			"%d. (%d,%d)\n",
			i+1,
			kartu[i].Sisi1,
			kartu[i].Sisi2,
		)
	}
}

func main() {

	dominoes := buatDominoes()

	kocokKartu(&dominoes)

	kartuPemain := make([]Domino, 7)

	for i := 0; i < 7; i++ {
		kartuPemain[i] = ambilKartu(&dominoes)
	}

	kartuMeja := ambilKartu(&dominoes)

	ujung := kartuMeja.Sisi2

	fmt.Println("===== PERMAINAN GAPLEH =====")

	fmt.Printf(
		"Kartu Awal Meja : (%d,%d)\n",
		kartuMeja.Sisi1,
		kartuMeja.Sisi2,
	)

	for {

		if len(kartuPemain) == 0 {

			fmt.Println(
				"\nSelamat! Semua kartu berhasil dimainkan.",
			)

			break
		}

		fmt.Println("\n====================")
		fmt.Println("Angka Ujung :", ujung)
		fmt.Println("Sisa Tumpukan :", dominoes.Jumlah)

		tampilkanKartu(kartuPemain)

		fmt.Println("\nMenu")
		fmt.Println("1. Mainkan Kartu")
		fmt.Println("2. Ambil Kartu")
		fmt.Println("0. Keluar")

		var menu int

		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		switch menu {

		case 1:

			var pilihan int

			fmt.Print(
				"Pilih nomor kartu yang ingin dimainkan : ",
			)

			fmt.Scan(&pilihan)

			if pilihan < 1 ||
				pilihan > len(kartuPemain) {

				fmt.Println("Pilihan tidak valid.")
				continue
			}

			kartu := kartuPemain[pilihan-1]

			if !bisaMain(kartu, ujung) {

				fmt.Println(
					"Kartu tidak bisa disambung.",
				)

				continue
			}

			fmt.Printf(
				"Kartu (%d,%d) dimainkan\n",
				kartu.Sisi1,
				kartu.Sisi2,
			)

			if kartu.Sisi1 == ujung {
				ujung = kartu.Sisi2
			} else {
				ujung = kartu.Sisi1
			}

			kartuPemain =
				hapusKartu(
					kartuPemain,
					pilihan-1,
				)

		case 2:

			if dominoes.Jumlah == 0 {

				fmt.Println(
					"Tumpukan kartu sudah habis.",
				)

				continue
			}

			kartuBaru :=
				ambilKartu(&dominoes)

			fmt.Printf(
				"Anda mengambil kartu (%d,%d)\n",
				kartuBaru.Sisi1,
				kartuBaru.Sisi2,
			)

			kartuPemain =
				append(
					kartuPemain,
					kartuBaru,
				)

		case 0:

			fmt.Println(
				"Permainan selesai.",
			)

			return

		default:

			fmt.Println(
				"Menu tidak valid.",
			)
		}
	}
}