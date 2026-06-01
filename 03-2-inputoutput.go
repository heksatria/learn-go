package main

import "fmt"

func latIO_output(){
	nama := "Budi"
	umur := 20

	// Println — print + newline
	fmt.Println("Hello!")
	fmt.Println("Nama:", nama, "| Umur:", umur)

	// Printf — print dengan format
	fmt.Printf("Nama: %s, Umur: %d tahun\n", nama, umur)
	
	// Sprintf — format tapi simpan ke string
	pesan := fmt.Sprintf("Halo, %s!", nama)
	fmt.Println(pesan)
}
