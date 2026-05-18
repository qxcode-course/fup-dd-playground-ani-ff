package main

import "fmt"

func main() {
    var B, T int
    fmt.Scan(&B)
    fmt.Scan(&T)

    soma := B + T

    if soma > 160 {
        fmt.Println(1)
    } else if soma < 160 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}