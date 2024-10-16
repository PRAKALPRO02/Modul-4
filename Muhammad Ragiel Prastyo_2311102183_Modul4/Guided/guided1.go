//MUHAMMAD RAGIEL PRASTYO
//2311102183
package main
import "fmt"

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Prosedur untuk menghitung dan menampilkan permutasi
func permutasi(n, r int) {
	hasilPermutasi := factorial(n) / factorial(n-r)
	fmt.Printf("Permutasi dari %dP%d adalah: %d\n", n, r, hasilPermutasi)
}

func main() {
	// Memanggil prosedur untuk menghitung dan menampilkan permutasi
	n, r := 5, 3
	permutasi(n, r)
}