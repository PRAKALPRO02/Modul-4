package main

import "fmt"

func deret(n int) {
	var hasil int = n
	fmt.Print(hasil, " ")
	for {
		if hasil%2 == 0 {
			hasil = hasil / 2
			fmt.Print(hasil, " ")
		} else if hasil%2 == 1 {
			hasil = 3*hasil + 1
			fmt.Print(hasil, " ")
		}
		if hasil == 1 {
			break
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	deret(n)
}