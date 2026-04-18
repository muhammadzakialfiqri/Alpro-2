package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	cx, cy, r int
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.cx, &c1.cy, &c1.r)
	fmt.Scan(&c2.cx, &c2.cy, &c2.r)
	fmt.Scan(&p.x, &p.y)

	d1 := didalam(c1, p)
	d2 := didalam(c2, p)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func didalam(c lingkaran, p titik) bool {
	return jarak(p, titik{c.cx, c.cy}) <= float64(c.r)
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}