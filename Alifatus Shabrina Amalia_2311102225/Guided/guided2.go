package main

import "fmt"

//deklarasi prosedur
func greet(name string) {
	fmt.Println("Hallo", name)
}

func main() {
	//manggil prosedur
	greet("Alifatus")
	greet("Shabrina")
}
