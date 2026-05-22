package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    for; a <= b; a++ {
        if a%15 == 0 {
            fmt.Println("zigzag")
        } else if a%3 == 0 {
            fmt.Println("zig")
        } else if a%5 == 0 {
            fmt.Println("zag")
        } else {
            fmt.Println(a)
        }
    }
}
