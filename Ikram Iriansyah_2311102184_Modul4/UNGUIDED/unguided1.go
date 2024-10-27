package main

import "fmt"

func cetakDeret(n int) {

	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Input Nilai : ")
	fmt.Scan(&n)

	if n < 1000000 {
		fmt.Print("Suku Deret : ")
		cetakDeret(n)
	} else {
		fmt.Println("Inputan harus antara 1 sampai 1,000,000!")
	}
}
