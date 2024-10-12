package main

import "fmt"

func main() {
	var nilaiAwal int
	fmt.Print("Masukkan bilangan awal (integer positif di bawah 1 juta): ")
	fmt.Scan(&nilaiAwal)

	if validInput(nilaiAwal) {
		tampilkanDeret(nilaiAwal)
	} else {
		fmt.Println("Input tidak valid. Masukkan angka positif kurang dari 1 juta.")
	}
}

func validInput(x int) bool {
	return x > 0 && x < 1000000
}

func tampilkanDeret(m int) {
	for m > 1 {
		fmt.Printf("%d ", m)
		m = hitungBerikutnya(m)
	}
	fmt.Println(1) // Terakhir pasti 1
}

func hitungBerikutnya(bil int) int {
	if bil%2 == 0 {
		return bil / 2
	}
	return 3*bil + 1
}
