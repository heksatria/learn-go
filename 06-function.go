package main

import "fmt"

// function with no parameter and no return value
    func sayHello() {
        fmt.Println("Hello")
    }

    // function with parameter and no return value
    func sayHelloTo(name string) {
        fmt.Println("Hello", name)
    }

    // function with parameter and return value
    func add(a, b int) int {
        return a + b
    }

    // function with multiple return value
    func divide(a, b float32) (float32, int) {
        return a / b, int(a) % int(b)
    }

    // function with named return value
    func getCoordinates(kjg,  lb int) int {
        keliling:=2*kjg+2*lb 
        return keliling
    }

    // Multiplede return value with named return value
    func getName()(string,string){
        return "Hendri","Ekasatria"
    }

    func persegi(pjg, lbr int) string {
        keliling := 2*pjg + 2*lbr
        luas := pjg * lbr
        return fmt.Sprintf("Keliling: %d, Luas: %d", keliling, luas)
    }
    func persegiPanjang(pjg, lbr int) (kll int,luas int,hasil string) {
        kll = 2*pjg + 2*lbr
        luas = pjg * lbr
        hasil = fmt.Sprintf("Keliling: %d, Luas: %d", kll, luas)
        return
    }

    //Variadic function
    func sumAll(numbers ...int) int { //...int -->seolah sebuah slice
        total := 0
        for _, number := range numbers {
            total += number
        }
        return total
    }

    func function06(){
        total := sumAll(1, 2, 3, 4, 5)
        fmt.Println("Total:", total)
    }
    
    //Function as value
    func getGoodbye(name string) string{
        return "Goodbye, " + name
    }
    func function06b(){
        goodbye:= getGoodbye
        fmt.Println(goodbye("Hendri"))
    }

    //Function as parameter
    func sayHelloWithFilter(name string, filter func(string) string)  {
        filteredName := filter(name)
        fmt.Println( "Hello, " + filteredName)
    }
    //implementasi filter function
    func spamFilter(name string) string {
        if name == "anjing" {
            return "..."
        }
        return name
    }

    func function06c(){
        spam := spamFilter
        sayHelloWithFilter("Hendri", spam)
        sayHelloWithFilter("anjing", spam)
    }