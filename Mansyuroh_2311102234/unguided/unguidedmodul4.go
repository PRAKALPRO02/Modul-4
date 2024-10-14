package main

import "fmt"

func cetakDeret(n int) {
	fmt.Print(n)

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(" ", n)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan suku awal: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cetakDeret(n)

	fmt.Println()
}
