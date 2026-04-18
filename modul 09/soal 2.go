package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func main() {
	var arr [NMAX]int
	var n, x, del, cari int
	var sum, varian, rata, std float64
	var freq int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Semua elemen:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("Elemen indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("Elemen indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	fmt.Println("Elemen indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&del)
	for i := del; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah penghapusan:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	rata = sum / float64(n)
	fmt.Println("Rata-rata:", rata)

	for i := 0; i < n; i++ {
		varian += (float64(arr[i]) - rata) * (float64(arr[i]) - rata)
	}
	std = math.Sqrt(varian / float64(n))
	fmt.Println("Standar deviasi:", std)

	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}