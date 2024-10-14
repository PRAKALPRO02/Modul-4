// Fajar Farizqi Azmi
// 2311102192

package main

import "fmt"


func cetakDeret(n int) {
    for n != 1 {
        fmt.Print(n, "-" )
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 32*n + 1
        }
    }
    fmt.Println(1)
}

func main() {
    var n int
    fmt.Println("Masukkan bilangan bulat positif kurang dari 1000000: ") 
    fmt.Scan(&n)
    
    if n > 0 && n < 1000000 {
        cetakDeret(n)
    } else {
        fmt.Println("Masukkan bilangan positif kurang dari 1000000!") 
    }
}