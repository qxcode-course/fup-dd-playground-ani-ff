 package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    if A == B{
        fmt.Println("oi")
    } else if A > B {
        fmt.Println("invalido")
    }
}
