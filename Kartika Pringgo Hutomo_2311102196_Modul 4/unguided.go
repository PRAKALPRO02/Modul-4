package main

import "fmt"

func cetakDeret(n int) {
	hasil := n
	fmt.Print(hasil)

	for hasil != 1 {
		if hasil%2 == 0 {
			hasil = hasil / 2
		} else {
			hasil = 3*hasil + 1
		}
		fmt.Print(" ", hasil) // Tambahkan spasi sebelum angka berikutnya
	}
	fmt.Println() // Tambahkan baris baru setelah selesai mencetak deret
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	cetakDeret(n)
}
