package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	if n < 1 || n >= 1000000 {
		fmt.Println("Bilangan harus lebih besar dari 0 dan kurang dari 1000000")
		return
	}
	cetakDeret(n)
}
