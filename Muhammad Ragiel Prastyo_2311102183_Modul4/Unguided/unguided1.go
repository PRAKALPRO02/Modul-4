//MUHAMMAD RAGIEL PRASTYO
//2311102183
package main
import (
	"fmt"
)

// Fungsi cetakDeret untuk mencetak deret berdasarkan aturan yang diberikan
func cetakDeret(n int) {
	var deret []int
	
	for n != 1 {
		deret = append(deret, n) 
		if n%2 == 0 {           
			n = n / 2          
		} else {                 
			n = 3*n + 1         
		}
	}
	
	deret = append(deret, 1)
	
	// Mencetak deret
	for i, v := range deret {
		if i == len(deret)-1 {
			fmt.Print(v) 
		} else {
			fmt.Print(v, " ") 
		}
	}
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000000 {
		fmt.Println("Masukan tidak valid. Masukkan bilangan positif kurang dari 1000000.")
		return
	}

	cetakDeret(n)
}