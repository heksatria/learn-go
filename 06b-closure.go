package main

import "fmt"

func myclosure(){
	counter :=0
	increment:=func()int{
		counter++
		return counter
	}
	increment()
	increment()
	increment()
	fmt.Println(increment())
}
func showClosure(){
	myclosure()
	myclosure()
	myclosure()
}