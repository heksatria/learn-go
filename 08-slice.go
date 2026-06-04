package main
import "fmt"

func latSlice(){
	//Membuat slice dari array
	var buah = [...]string{"jambu","apel","pepaya","klengkeng","rambutan"}

	//Panjang array
	fmt.Println(len(buah)) //5

	var slice1 = buah[1:3]
	fmt.Println(slice1)  //[apel pepaya]

	//Slice adalah pointer ke array, buktinya
	slice1[0]= "pir"
	fmt.Println(slice1)  //[pir pepaya]
	
	//Menambah isi slice
	// append(slice1,"kelapa") //(value of type []string) is not used --> panjang array cuma 5
	//Solusi => buat slice baru
	var slice2 = append(slice1,"kelapa")
	fmt.Println(slice2)

	slice2 = append(slice2,"kelapa2")
	slice2 = append(slice2,"kelapa3")
	slice2 = append(slice2,"kelapa4")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))


}