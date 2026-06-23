package main

import "fmt"

//For klasik
func latLoop1() {
	// Loop For
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
}

// For pengganri while
func latLoop2(){
	var i int = 0
	for i <10{
		fmt.Println("Perulangan ke-", i)
		i++
	}
}

// Break pada loop
func latLoop3(){
	var i int = 1
	for {
		fmt.Println("Perulangan ke-", i)
		if(i == 200){
			break
		}
		i++
	}
}

func latLoop4(){
	fruits := []string{"nangka","pir","jambu"}
	for index,fruit :=range fruits{
		fmt.Println("Index ",index," =",fruit)
	}
}
func latLoop5(){
	fruits := []string{"nangka","pir","jambu"}
	for _,fruit :=range fruits{
		fmt.Println("Buah ",fruit)
	}
}
