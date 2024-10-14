package main

import "fmt"

func greet(name string) {
	fmt.Println("Hallo ", name)
}
func main() {
	name1 := "Ervan"
	name2 := "pandia"
	greet(name1)
	greet(name2)
}
