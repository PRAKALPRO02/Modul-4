package main
import "fmt"

func greet(name string) {
	fmt.Println("Hallo ", name)
}

func main() {
	name1 := "ADITHANA"
	name2 := "GANTENK"
	greet(name1)
	greet(name2)
}