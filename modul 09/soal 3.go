package main

import "fmt"

const NMAX = 100

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [NMAX]string
	var n int

	fmt.Scan(&klubA)
	fmt.Scan(&klubB)

	for {
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}
		n++
	}

	for i := 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}