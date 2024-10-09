package main

import "fmt"

func cetakDeret(panji int) {
    for panji != 1 {
        fmt.Print(panji, " ")
        if panji%2 == 0 {
            panji = panji / 2
        } else {
            panji = 3*panji + 1
        }
    }
    fmt.Println(panji)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai awal: ")
    fmt.Scanln(&n)
    cetakDeret(n)
}
