package main

import "fmt"

func menuMenampilkanData() {
	var opsi string
	fmt.Println("[]============================================[]")
	fmt.Println("[]               Tampilkan Data               []")
	fmt.Println("[]============================================[]")
	fmt.Println("[]====[]======================================[]")
	fmt.Println("[] 1. [] Tampilkan Data Perusahaan            []")
	fmt.Println("[]====[]======================================[]")
	fmt.Println("[] 2. [] Tampilkan Data Karyawan              []")
	fmt.Println("[]====[]======================================[]")
	fmt.Println("[] 0. [] Kembali                              []")
	fmt.Println("[]====[]======================================[]")
	fmt.Print("Masukkan inputan (0/2): ")
	fmt.Scan(&opsi)
	
	for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "0" {
		fmt.Println("Masukkan inputan yang valid (0/4)")
		fmt.Scan(&opsi)
	}
	if opsi == "1" {
		menampilkanDataPer(&P)
	} else if opsi == "2" {
		menampilkanData(K, n)
	} else if opsi == "3" {
		selectionSort(&K, n)
		menampilkanData(K, n)
	}  else if opsi == "0" {
		menu()
	}
}