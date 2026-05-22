package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i := a; i < b; i++ {
        if i > a {
            fmt.Print(" ")
        }
        fmt.Print(i)
    }
    fmt.Println(" ]")
}
