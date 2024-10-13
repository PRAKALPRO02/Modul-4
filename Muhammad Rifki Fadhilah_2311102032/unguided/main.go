package main

import "fmt"

func deretBilangan(n int){
  if n < 1000{
    for n != 1{
      fmt.Print(n," ")
      if n % 2 == 0 {
        n = n/ 2
      }else{
        n = 3*n+ 1
      }
    }
    fmt.Print(n)
  }else{
    fmt.Print("Bilangan lebih dari 1000")
  }
}



// Prosedur untuk mencetak pesan
func cetakPesan(pesan string) {
  fmt.Println(pesan)
}

func main() {
  // Memanggil prosedur
  cetakPesan("Halo, Fadhil")
}

