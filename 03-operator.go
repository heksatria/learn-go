package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func latOperator() {
	var(
		panjang int = 0
		lebar int = 0
	)

	fmt.Print("Nama area :")
	reader := bufio.NewReader(os.Stdin)
	area, _ := reader.ReadString('\n')
	area = strings.TrimSpace(area)
	fmt.Print("Masukan panjang :"); fmt.Scan(&panjang);
	fmt.Print("Masukan lebar :"); fmt.Scan(&lebar);

	luas := panjang * lebar
	fmt.Println("Luas ", area, luas)
	fmt.Println(unsafe.Sizeof(luas))
}

func latJenisOperator() {
		a, b := 10, 3

	// Operator Aritmatika
	fmt.Println(a + b) // 13 — penjumlahan
	fmt.Println(a - b) // 7 — pengurangan
	fmt.Println(a * b) // 30 — perkalian
	fmt.Println(a / b) // 3 — pembagian (hasil integer)
	fmt.Println(a % b) // 1 — modulus (sisa bagi)

	// Operator Perbandingan
	fmt.Println(5 == 5) // true
	fmt.Println(5 != 3) // true
	fmt.Println(5 > 3) // true
	fmt.Println(5 < 3) // false
	fmt.Println(5 >= 5) // true
	fmt.Println(5 <= 4) // false

	// Operator Logika
	fmt.Println(true && false) // false — AND
	fmt.Println(true || false) // true — OR
	fmt.Println(!true) // false — NOT

	// Operator Penugasan
	x := 10
	x += 5 // x = x + 5 → 15
	x -= 3 // x = x - 3 → 12
	x *= 2 // x = x * 2 → 24
	x /= 4 // x = x / 4 → 6
	x++ // x = x + 1 → 7
	x-- // x = x - 1 → 6
}