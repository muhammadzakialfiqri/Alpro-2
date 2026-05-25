	package main

	import "fmt"

	const nMax int = 7919

	type Buku struct {
		id, judul, penulis, penerbit string
		eksemplar, tahun, rating     int
	}

	type DaftarBuku [nMax]Buku

	func DaftarkanBuku(pustaka *DaftarBuku, n int) {
		var i int

		for i = 0; i < n; i++ {

			fmt.Scan(
				&pustaka[i].id,
				&pustaka[i].judul,
				&pustaka[i].penulis,
				&pustaka[i].penerbit,
				&pustaka[i].eksemplar,
				&pustaka[i].tahun,
				&pustaka[i].rating,
			)

		}
	}

	func CetakTerfavorit(pustaka DaftarBuku, n int) {
		var iMax, i int

		iMax = 0

		for i = 1; i < n; i++ {

			if pustaka[i].rating > pustaka[iMax].rating {
				iMax = i
			}

		}

		fmt.Println()
		fmt.Println("Buku Terfavorit")
		fmt.Println("Judul :", pustaka[iMax].judul)
		fmt.Println("Penulis :", pustaka[iMax].penulis)
		fmt.Println("Penerbit :", pustaka[iMax].penerbit)
		fmt.Println("Tahun :", pustaka[iMax].tahun)
	}

	func UrutBuku(pustaka *DaftarBuku, n int) {
		var i, j int
		var temp Buku

		for i = 1; i < n; i++ {

			temp = pustaka[i]
			j = i

			for j > 0 && temp.rating > pustaka[j-1].rating {

				pustaka[j] = pustaka[j-1]
				j--

			}

			pustaka[j] = temp
		}
	}

	func Cetak5Terbaru(pustaka DaftarBuku, n int) {
		var i, batas int

		fmt.Println("5 Rating Tertinggi")

		if n < 5 {
			batas = n
		} else {
			batas = 5
		}

		for i = 0; i < batas; i++ {
			fmt.Println(pustaka[i].judul)
		}
	}

	func CariBuku(pustaka DaftarBuku, n int, r int) {
		var kiri, kanan, tengah int
		var ketemu bool

		kiri = 0
		kanan = n - 1

		for kiri <= kanan && !ketemu {

			tengah = (kiri + kanan) / 2

			if pustaka[tengah].rating == r {

				ketemu = true

				fmt.Println()
				fmt.Println("Buku Ditemukan")
				fmt.Println("Judul :", pustaka[tengah].judul)
				fmt.Println("Penulis :", pustaka[tengah].penulis)
				fmt.Println("Penerbit :", pustaka[tengah].penerbit)
				fmt.Println("Tahun :", pustaka[tengah].tahun)
				fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
				fmt.Println("Rating :", pustaka[tengah].rating)

			} else if r > pustaka[tengah].rating {

				kanan = tengah - 1

			} else {

				kiri = tengah + 1

			}
		}

		if !ketemu {
			fmt.Println("Tidak ada buku dengan rating seperti itu")
		}
	}

	func main() {

		var pustaka DaftarBuku
		var n, ratingCari int

		fmt.Scan(&n)

		DaftarkanBuku(&pustaka, n)

		fmt.Scan(&ratingCari)

		CetakTerfavorit(pustaka, n)

		UrutBuku(&pustaka, n)

		Cetak5Terbaru(pustaka, n)

		CariBuku(pustaka, n, ratingCari)
	}