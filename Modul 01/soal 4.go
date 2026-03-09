package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin float32

	fmt.Print("masukan nilai suhu dalam celcius : ")
	fmt.Scan(&celcius)

	fahrenheit = (celcius + 32) * 9 / 5
	reamur = celcius * 4 / 5
	kelvin = celcius + 273
	fmt.Printf("fahrenheit = %.0f\n", fahrenheit)
	fmt.Printf("reamur = %.0f\n", reamur)
	fmt.Printf("kelvin = %.0f\n", kelvin)
}
