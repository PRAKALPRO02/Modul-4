package main 

import "fmt"

func faktorial(n int ) int {
	if n == 0 || n == 1{
		return 1
	}

	result := 1
	for i := 2; i <= n; i++{
		result *= i
	}
	return result
}

func permutasi(n, r int) int{
	return faktorial(n) / faktorial(n-r)

}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main(){
	var a, b, c, d int 

	fmt.Println()
	fmt.Print("Masukkan bilangan a, b, c, d (dengan spasi) : ")

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	
	/*
	syarat :
	- a >= c (baris pertama hasil permutasi dan kombinasi a terhadap c)
	- b >= d (baris kedua hasil permutasi dan kombinasi b terhadap d)
	
	
	*/
	if a >= c && b >= d {

		// jika syarat terpenuhi hitung permutasi dan kombinasi a c
		permutasiAC := permutasi(a, c)
		kombinasiAC := kombinasi(a, c)

		permutasiBD := permutasi(b, d)
		kombinasiBD := kombinasi(b, d)


		fmt.Println("Permutasi (a,c) dan kombinasi (a,c) : ", permutasiAC, kombinasiAC)
		fmt.Println("Permutasi (b,d) dan kombinasi (b,d) : ", permutasiBD, kombinasiBD)
	} else {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
	}
	fmt.Println()
}

