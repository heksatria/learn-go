package main
import "fmt"

//Slice dari array
func latSlice1(){
	//Membuat slice dari array
	var buah = [...]string{"jambu","apel","pepaya","klengkeng","rambutan"}

	//Panjang array
	fmt.Println("Panjang array buah : ",len(buah)) //5

	var slice1 = buah[1:3]
	fmt.Println("Panjang slice-1 : ",len(slice1))  //[apel pepaya]
	fmt.Println("Kapasitas slice-1 : ", cap(slice1))  // apel sampai index array buah terakhir
	fmt.Println("Isi slice-1 : ",slice1)  //[apel pepaya]

	//Slice adalah pointer ke array, buktinya
	slice1[0]= "pir"
	fmt.Println("Isi array buah sekarang :  ",buah)  //[pir pepaya]
	
	//Menambah isi slice
	// append(slice1,"kelapa") //(value of type []string) is not used --> panjang array cuma 5
	//Solusi => buat slice baru
	var slice2 = append(slice1,"kelapa")
	fmt.Println("Kapasitas slice-2 : ",cap(slice2))
	fmt.Println("Isi slice-2 : ",slice2)

	slice2 = append(slice2,"kelapa2")
	slice2 = append(slice2,"kelapa3")
	slice2 = append(slice2,"kelapa4")
	slice2 = append(slice2,"kelapa5")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	slice2 = append(slice2,"kelapa6") //Otomatis kapasitas slice ditambah jadi 2 kali lipat
	slice2 = append(slice2,"kelapa7")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))


}

// Slice 
func latSlice2(){
	var nilai =[]int{80,75,66,90,87}
	
	fmt.Println("Panjang slice : ",len(nilai))
	fmt.Println("Kapasitas slice : ",cap(nilai))
	fmt.Println("Isi slice : ",nilai)

	// append.nilai(99) // error karena kapasitas cuma 5
}

//Create slice using make
func latSlice3(){
	var sayur = make([]string,3,5) //panjang =3, kapasitas =5 
	fmt.Println(len(sayur)) //3
	fmt.Println(cap(sayur)) //5
	fmt.Println(sayur) //--> ["", "", ""]

	sayur[0] = "bayam" //["bayam", "", ""]
	sayur = append(sayur,"kangkung")

	fmt.Println(len(sayur)) //4
	fmt.Println(cap(sayur)) //5
	fmt.Println(sayur)  //["bayam", "", "", "kangkung"]

	sayur = append(sayur,"jamur") 
	fmt.Println(len(sayur)) //5
	fmt.Println(cap(sayur)) //5
	fmt.Println(sayur) //["bayam", "", "", "kangkung","jamur"]

	sayur = append(sayur,"nangka") 
	fmt.Println(len(sayur)) //6
	fmt.Println(cap(sayur)) //10
	fmt.Println(sayur) //["bayam", "", "", "kangkung","jamur","nangka"]
}