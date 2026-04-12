package main
import "fmt"

func bintang(baris, n int){

	if baris > n{
		return
	}

	for i := 0 ; i < baris ; i++{
		fmt.Print("*")
	}
	fmt.Println()

	bintang(baris + 1, n )
}

func main(){

	var n int
	fmt.Scan(&n)

	bintang(1, n)
}