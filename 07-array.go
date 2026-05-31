package main

import "fmt"

func latArray() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array:", arr)
	fmt.Println("Panjang array:", len(arr))
}

func latArray2() {
	var buah =[5]string{"apel", "jeruk", "mangga", "pisang", "anggur"}
	fmt.Println("Panjang array:", len(buah))
	fmt.Println("Buah:", buah[0:3])
	fmt.Println("Buah:", buah[3:])
	fmt.Println("Buah:", buah[1:3])

}