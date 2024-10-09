package main
//2311102038
import "fmt"

func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }//2311102038
    }
    fmt.Println(n) 
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")//2311102038
    fmt.Scan(&n)
//2311102038
    cetakDeret(n)
}//2311102038
//2311102038