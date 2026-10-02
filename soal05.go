package main

import "fmt"

func main() {
	var gajiPokok, jamLembur int

	// Input
	fmt.Scan(&gajiPokok, &jamLembur)

	// Bonus lembur
	bonusLembur := 45000 * jamLembur

	// Potongan asuransi 2% + dana pensiun 3,5%
	potongan := (2*gajiPokok)/100 + (35*gajiPokok)/1000

	// Gaji bersih
	gajiBersih := gajiPokok + bonusLembur - potongan

	// Output
	fmt.Println(gajiBersih)
}