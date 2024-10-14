package main

import "fmt"

func printBaris(n int) {

	for n != 1 {
		fmt.Printf("%d ", n)
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
	fmt.Print("Input (n): ")
	fmt.Scan(&n)

	if n > 0 && n < 1000000 {
		printBaris(n)
	} else {
		fmt.Println("Tidak Memenuhi kondisi 'n > 0 && n < 1000000'")
	}
}
