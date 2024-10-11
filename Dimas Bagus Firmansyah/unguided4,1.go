package main

import "fmt"

// Fungsi untuk mencetak deret sesuai aturan yang diberikan
func cetakDeret(n int) {
    for n != 1 {
        fmt.Print(n, " ")
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }
    fmt.Println(1) // 2311102002
}

func main() {
    var n int
    fmt.Println("Masukkan bilangan bulat positif kurang dari 1000000: ") // 2311102002
    fmt.Scan(&n)
    
    if n > 0 && n < 1000000 {
        cetakDeret(n)
    } else {
        fmt.Println("Masukkan bilangan positif kurang dari 1000000!") // 2311102002
    }
}