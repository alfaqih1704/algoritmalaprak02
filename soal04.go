package main

import "fmt"

func main() {
	var emas, sisa, perak, tembaga, koin int
	fmt.Scan(&koin)
	emas = koin / 9
	sisa = koin % 9
	perak = sisa / 3
	tembaga = koin % 3

	fmt.Println(emas, perak, tembaga)
}