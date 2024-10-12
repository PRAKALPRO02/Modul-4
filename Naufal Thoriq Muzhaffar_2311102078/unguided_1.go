package main
import "fmt"

// Prosedur Deret
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Printf("%d\n", n) // Cetak angka 1 di akhir
}

func main() {
	// Masukkan nilai suku awal
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Panggil prosedur cetakDeret dengan nilai n
	cetakDeret(n)
}
