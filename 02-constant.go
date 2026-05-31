package main

import "fmt"

func latConst() {
	const(
		PI float64 = 3.14159
		E  float64 = 2.71828
		status string = "aktif"
		
	)
	fmt.Println("PI:", PI)
	fmt.Println("E:", E)
	fmt.Println("Status:", status)
}