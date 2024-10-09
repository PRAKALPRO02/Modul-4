package main

import "fmt"

func greet(name string) {
	fmt.Println("Hallo ", name)
}
func main() {
	name1 := "ahmadan"
	name2 := "dimas"
	name3 := "agha"
	greet(name1)
	greet(name2)
	greet(name3)
}
