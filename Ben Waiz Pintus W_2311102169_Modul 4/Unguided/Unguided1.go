package main

import "fmt"

func hasilkanDeret(bilangan int) {
	for bilangan != 1 {
		fmt.Printf("%d ", bilangan)

		if bilangan%2 == 0 {
			bilangan /= 2
		} else {
			bilangan = 3*bilangan + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)
	if bilangan > 1000000 {
		fmt.Println("Bilangan harus kurang dari 1000000 !!!")
	} else {
		hasilkanDeret(bilangan)
	}
}
