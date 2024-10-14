package main

import "fmt"

func urutan(n int) {
	kondisi := true
	for kondisi == true {
		if n != 1 {
			kondisi = true
			fmt.Print(n, " ")
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
		} else {
			fmt.Print(n, " ")
			kondisi = false
		}

	}
}

func main() {
	var n int
	fmt.Scan(&n)
	urutan(n)
}
