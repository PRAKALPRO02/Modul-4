package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 && n < 1000000 {
		PrintDeret_219(n)
	}
}

func PrintDeret_219(n int) {
	fmt.Print(n, " ")
	if n%2 == 0 { // pengecekan bilangan genap
		PrintDeret_219(n / 2)
	} else if n != 1 { // pengecekan bilangan ganjil
		PrintDeret_219(3*n + 1)
	}
}
