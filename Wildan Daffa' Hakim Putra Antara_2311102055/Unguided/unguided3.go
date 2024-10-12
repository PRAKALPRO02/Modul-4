package main

import "fmt"

func cetakDeret(n int) {
	fmt.Print(n, " ")
	if n%2 == 0 {
		cetakDeret(n / 2)
	} else if n != 1 {
		cetakDeret(3*n + 1)
	} else {
		// kondisi else ini bisa diremove namun saat cetak akan meninggalkan "%" 
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n > 1000000 {
		fmt.Println("Bilangan harus <= 1000000")
	} else {
		cetakDeret(n)
	}
}
