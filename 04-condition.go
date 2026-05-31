package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func latIF() {
    //Grading
	nilai := 0

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 90 {
		fmt.Println("Grade: A")
		} else if nilai >= 80 {
		fmt.Println("Grade: B")
		} else if nilai >= 70 {
		fmt.Println("Grade: C")
		} else if nilai >= 60 {
		fmt.Println("Grade: D")
		} else {
		fmt.Println("Grade: E")
	}
}

func hitungSkor() int {
	return 60
}

// Latihan IF dengan short statement
/* if init_statement; condition {
     ...
 }
 */
func latIFShortStt() {
	if skor := hitungSkor(); skor >= 70 {
		fmt.Println("Lulus!")
	} else {
		fmt.Printf("Tidak lulus. Skor: %d\n", skor)
	}
}

// Latihan Switch
func latSwitch(){
	hari := ""

	fmt.Print("Masukan hari : ")
	fmt.Scan(&hari)

	switch(strings.ToLower(hari)){
		case "senin","selasa","rabu","kamis","jumat":
			fmt.Println(hari," adalah hari kerja")
		case "sabtu","minggu":
			fmt.Println(hari," adalah hari libur")
		default:
			fmt.Println("Input Anda salah")

	}

}
func latSwitch2(){
	nilai:=0
	fmt.Print("Masukan nilai : ")
	fmt.Scan(&nilai)

	switch{
	case nilai >=90:
		fmt.Println("Grade : A")
	case nilai >=80:
		fmt.Println("Grade : B")
	case nilai >=70:
		fmt.Println("Grade : C")
	case nilai >=60:
		fmt.Println("Grade : D")
	default:
		fmt.Println("Grade : E")
	}
}

// Latihan Falltrough pada Switch
func latSwitch3(){
	statusHari :=""
	fmt.Println("Masukan status hari (hari kerja/hari libur) :")
	reader := bufio.NewReader(os.Stdin)
	statusHari,_=reader.ReadString('\n')
	statusHari = strings.TrimSpace(statusHari)

	switch{
	case strings.ToLower(statusHari) == "hari kerja" || strings.ToLower(statusHari) == "kerja":
		fmt.Println("Senin")
		fallthrough
	case strings.ToLower(statusHari) == "hari kerja" || strings.ToLower(statusHari) == "kerja":
		fmt.Println("Selasa")
		fallthrough
	case strings.ToLower(statusHari) == "hari kerja" || strings.ToLower(statusHari) == "kerja":
		fmt.Println("Rabu")
		fallthrough
	case strings.ToLower(statusHari) == "hari kerja" || strings.ToLower(statusHari) == "kerja":
		fmt.Println("Kamis")
		fallthrough
	case strings.ToLower(statusHari) == "hari kerja" || strings.ToLower(statusHari) == "kerja":
		fmt.Println("Jumat")
	case strings.ToLower(statusHari) == "hari libur" || strings.ToLower(statusHari) == "libur":
		fmt.Println("Sabtu")
		fallthrough
	case strings.ToLower(statusHari) == "hari libur" || strings.ToLower(statusHari) == "libur":
		fmt.Println("Minggu")
	default:
		fmt.Println("Input Anda salah")	
	}
}