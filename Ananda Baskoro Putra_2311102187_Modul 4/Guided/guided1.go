package main

import "fmt"

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

func permutasi(n, r int) {
	hasilPermutasi := factorial(n) / factorial(n-r)
	fmt.Printf("Permutasi dari %dP%d adalah: %d\n", n, r, hasilPermutasi)
}

func main() {

	n, r := 6, 2
	permutasi(n, r)
}