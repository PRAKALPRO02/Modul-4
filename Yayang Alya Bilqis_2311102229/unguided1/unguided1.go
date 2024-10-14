package main

import "fmt"

func deretSkienaRekursif(n int) {
	fmt.Print(n, " ")
	if n != 1 {
		if n%2 == 0 {
			deretSkienaRekursif(n / 2)
		} else {
			deretSkienaRekursif(3*n + 1)
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan awal: ")
	fmt.Scan(&n)

	deretSkienaRekursif(n)
}
