package main

import "fmt"

// deklarasi prosedur
func greet(name string) {
	fmt.Println("Hallo", name)
}

func main() {
	// memanggil prosedur
	greet("uroh")
	greet("mansyur")
}
