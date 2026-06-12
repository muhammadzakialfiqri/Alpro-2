package main

import "fmt"

func main() {
	var n int
	var jumlah float64
	var si, sip1 float64

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		if i%2 == 1 {
			si = 1.0 / float64(2*i-1)
		} else {
			si = -1.0 / float64(2*i-1)
		}

		jumlah = jumlah + si

		if (i+1)%2 == 1 {
			sip1 = 1.0 / float64(2*(i+1)-1)
		} else {
			sip1 = -1.0 / float64(2*(i+1)-1)
		}

		selisih := si - sip1

		if selisih < 0 {
			selisih = -selisih
		}

		if selisih <= 0.00001 {
			fmt.Printf("Hasil PI: %.10f\n", 4*jumlah)
			fmt.Println("Pada i ke:", i)
			break
		}
	}
}
