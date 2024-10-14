package main

import "fmt"

func cetakDeret(n int, deret *[]int) {
	*deret = append(*deret, n)
	fmt.Print(n, " ")
	if n%2 == 0 {
		cetakDeret(n/2, deret)
	} else if n != 1 {
		cetakDeret(3*n+1, deret)
	} else {
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n > 1000000 {
		fmt.Println("Bilangan harus <= 1000000")
	} else {
		var deret []int
		cetakDeret(n, &deret)
		fmt.Println("Deret angka:")
		fmt.Println(deret)
	}
}
