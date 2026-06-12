package main

import "fmt"

func main() {
	var x, data string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	ada := false
	posisi := 0
	jumlah := 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&data)

		if data == x {
			jumlah++

			if !ada {
				ada = true
				posisi = i
			}
		}
	}

	fmt.Println()

	if ada {
		fmt.Println("a. Ada")
	} else {
		fmt.Println("a. Tidak Ada")
	}

	if ada {
		fmt.Println("b. Posisi:", posisi)
	} else {
		fmt.Println("b. Tidak ditemukan")
	}

	fmt.Println("c. Jumlah:", jumlah)

	if jumlah >= 2 {
		fmt.Println("d. Ya")
	} else {
		fmt.Println("d. Tidak")
	}
}