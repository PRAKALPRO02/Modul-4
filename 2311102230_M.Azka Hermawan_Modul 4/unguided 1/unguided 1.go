package main

import "fmt"

func printDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var bilAwal int
	fmt.Print("Masukan: ")
	fmt.Scan(&bilAwal)
	fmt.Print("Keluaran: ")
	printDeret(bilAwal)
}
