package main

import "fmt"

func greet(name string) {
	fmt.Println("Hallo ", name)
}
func main() {
	name1 := "panji"
	name2 := "panjay"
	greet(name1)
	greet(name2)
}