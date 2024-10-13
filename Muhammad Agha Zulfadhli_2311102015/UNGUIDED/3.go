package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	}
}

func cetakDeret(n int) {
	fmt.Print(n, " ")
	if n%2 == 0 {
		cetakDeret(n / 2)
	} else if n != 1 {
		cetakDeret(3*n + 1)
	}
}
