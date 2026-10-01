package main

import "fmt"

func main() {
	var N, Sisa, Jam, Menit, Detik int

	fmt.Scan(&N)
	Jam = N / 3600
	Sisa = N % 3600
	Menit = Sisa / 60
	Detik = Sisa % 60

	fmt.Println(Jam, Menit, Detik)
}