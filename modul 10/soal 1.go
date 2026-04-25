package main
import "fmt"



func main(){
	const Nmax = 1000
	var berat [Nmax] float64
	var n int
	

	fmt.Scan(&n)

	if n <= 0 || n > Nmax{
		return
	}

	for i := 0 ; i < n ; i++{
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1 ; i < n ; i++{
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Println()
	fmt.Print(min," ", max)
}