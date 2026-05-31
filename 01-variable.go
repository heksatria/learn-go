package main

import "fmt"

func latVariable(){
	var name string = "John Doe"
	var age int8 = 30
	var(
		country string = "USA"
		zipCode int16 = 12345
	)
	isProgrammer := true

	kabar := ""

	fmt.Println("Hello, " + name + "! You are " + fmt.Sprint(age) + " years old.")
	fmt.Println("You live in " + country + " with zip code " + fmt.Sprint(zipCode) + ".")
	fmt.Println("Are you a programmer? " + fmt.Sprint(isProgrammer))
	fmt.Print("Apa kabarmu bro?")
	fmt.Scan(&kabar)
	fmt.Println("Kabarmu:", kabar)
}

func latAlias(){
	type noKtp string

	var noKtpKoe noKtp = "0987654321"

	fmt.Println("No KTP Koe:", noKtpKoe)
}
