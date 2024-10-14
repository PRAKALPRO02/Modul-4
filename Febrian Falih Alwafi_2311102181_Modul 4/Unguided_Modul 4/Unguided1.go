package main

import "fmt"

// Fungsi untuk mencetak deret berdasarkan aturan tertentu
func cetakDeret(n int) {
	for {
		// Mencetak nilai n
		fmt.Printf("%d ", n)

		// Jika n mencapai 1, keluar dari loop
		if n == 1 {
			break
		}

		// Aturan deret: jika genap, bagi dua; jika ganjil, kalikan 3 dan tambahkan 1
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println() // Menambahkan baris baru setelah deret
} //Febrian Falih Alwafi
// 2311102181

func main() {
	var n int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&n)

	cetakDeret(n) // Memanggil fungsi untuk mencetak deret
}
