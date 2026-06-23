package main

import (
	"fmt"
	"strconv"
)

func latIO_output() {
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

func shortstt() {
	input := "abc"
	if angka, err := strconv.Atoi(input); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Angka:", angka)
	}
}
