package main
import "fmt"

func latSlice(){
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