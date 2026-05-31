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
    func getCoordinates() (x, y int) {
        x = 10
        y = 20
        return
    }
    
