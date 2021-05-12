package main


import (
    "fmt"
    "math"
)


func add(x, y int) int{
    return x + y
}

func main() {
    var test = "hello"
    fmt.Println(test + " world")

    fmt.Println("pi: ", math.Pi)

    res := add(3,5)
    fmt.Println("3 + 5 = ", res)
}

