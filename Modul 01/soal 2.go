package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("masukan tahun : ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Print("kabisat : true")
	} else {
		fmt.Print("kabisat : false")
	}
}
