package main

import "fmt"

func cetakDeret(n int) {
	fmt.Printf("%d\n", n)
	fmt.Printf("%d ", n)

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif kurang dari 1000000: ")
	fmt.Scan(&n)

	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukan tidak valid. Harap masukkan bilangan antara 1 hingga 999999.")
	}
}
