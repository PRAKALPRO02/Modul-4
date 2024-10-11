package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		n = hitungPanjangLangkah(n)
	}
	fmt.Println(n)
}

func hitungPanjangLangkah(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	var n int
	fmt.Print("Masukkan nilai suku awal (int positif kurang dari 1000000): ")
	fmt.Scan(&n)

	// Validasi input
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("nilai input tidak valid,Input harus bilangan positif dan kurang dari 1000000.")
	}
}
