package main

import "fmt"

func cetakDeret(n int) {
	var hasil int = n
	fmt.Print(hasil)
	for hasil != 1 {
		if hasil%2 == 0 {
			hasil = hasil / 2
			fmt.Print(hasil)
		} else {
			hasil = 3*hasil + 1
			fmt.Print(hasil)
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&n)

	cetakDeret(n)

}
