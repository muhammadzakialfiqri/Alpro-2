package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab, temp tabel
	var n int

	isiArray(&tab, &n)

	for i := 0; i < n; i++ {
		temp[i] = tab[i]
	}

	fmt.Println("Reverse teks :")
	balikanArray(&temp, n)
	cetakArray(temp, n)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrom(tab, n))
}