package main

import "fmt"

func cetakDeret(x *int) {
	fmt.Printf("%v ", *x)
	if *x > 1 {
		if *x%2 == 0 {
			*x /= 2
		} else {
			*x = *x*3 + 1
		}
		cetakDeret(x)
	}
}

func main() {
	var input int
	fmt.Print("Masukkan angka positif yang lebih kecil dari 1,000,000 : ")
	fmt.Scanln(&input)
	if input > 1000000 {
		fmt.Println("Angka lebih dari satu juta")
	} else {
		cetakDeret(&input)
	}
}
