package main

import "fmt"

func cetakSaldoCuti(K tabKaryawan, n int) {
	//print cuti
	fmt.Println("[]============================================================[]")
	fmt.Println("[]=================[Data Saldo Cuti Karyawan]=================[]")
	fmt.Println("[]============================================================[]")
	fmt.Println("[]<><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>[]")

	for i := 0; i < n; i++ {
		fmt.Println("[][Karyawan ke-", i+1, "]===========================================[]")
		fmt.Println("Nama		:", K[i].nama)
		fmt.Println("Saldo Cuti	:", K[i].cuti.saldo)
		fmt.Println("[]============================================================[]")
	}
	menuCuti(&K, n)
}
