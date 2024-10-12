package main

import "fmt"

func main() {
	var nilai int
	var teks string

	fmt.Scan(&nilai, &teks)

	tampilkanPesan(nilai, teks)
}
func tampilkanPesan(kode int, pesan string) {
	var kategori string

	switch kode {
	case 0:
		kategori = "kesalahan"
	case 1:
		kategori = "peringatan"
	case 2:
		kategori = "informasi"
	default:
		kategori = "tidak diketahui"
	}

	fmt.Printf("%s [%s]\n", pesan, kategori)
}
