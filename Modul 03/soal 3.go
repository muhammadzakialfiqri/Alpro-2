package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Print("masukan titik pusat lingkaran 1 dan 2 serta jari-jari dan titik sembarang : ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 := didalam(cx1, cy1, r1, x, y)
	d2 := didalam(cx2, cy2, r2, x, y)

	if d1 && d2 {
		fmt.Println("titik berada didalam lingkaran 1 dan 2")

	} else if d1 {
		fmt.Println("titik berada didalam lingkaran 1")

	} else if d2 {
		fmt.Println("titik berada didalam lingkaran 2")

	} else {
		fmt.Println("titik berada di luar lingkaran 1 dan 2")
	}
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))

}
