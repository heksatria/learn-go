package main

import (
	"fmt"
	"runtime"
)

func cekHari() {
	hari := "selasa"
	switch hari {
	case "senin":
		fmt.Println("Hari ini adalah hari senin")
	case "selasa":
		fmt.Println("Hari ini adalah hari selasa")
	case "rabu":
		fmt.Println("Hari ini adalah hari rabu")
	default:
		fmt.Println("Hari ini bukan senin, selasa, atau rabu")
	}
}

func cekOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
