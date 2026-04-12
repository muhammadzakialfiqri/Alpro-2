package main
import "fmt"

func faktor(n, angka int){

	if angka > n {
		return
	}

	if n % angka == 0 {
		fmt.Print(angka, " ")
	}

	faktor(n, angka + 1)

}

func main(){
	var n int

	fmt.Scan(&n)
	faktor(n, 1)
}