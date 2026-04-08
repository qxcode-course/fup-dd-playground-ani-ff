package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    mult := n % 7

    if mult == 0 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
