package main

import "fmt"

func cetakDeret (n int){
	fmt.Print(n)

	// loop hingga n != 1
	for n != 1 {
		if n%2 == 0 {
			// jika n genap, maka n dibagi dengan 2
			n = n/2
		} else {
			// jika n ganjil, n = 3n + 1
			n = 3*n + 1
		}
		fmt.Print(" ", n )
	}
	
}

func main(){

	var n int
	fmt.Printf("\nMasukan suku awal: ")
	fmt.Scan(&n)
	fmt.Print("Keluaran : ")
	cetakDeret(n)


	fmt.Println("\n\n")

}