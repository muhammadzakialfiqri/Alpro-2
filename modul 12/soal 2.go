package main

import "fmt"

func main() {

	var suara [21]int
	var x, suaraMasuk, suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {

		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i

		} else if suara[i] > suara[wakil] && i != ketua {
			wakil = i
		}
	}

	if suara[wakil] > suara[ketua] {
		ketua, wakil = wakil, ketua
	}
	fmt.Println()
	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
