package main
import "fmt"
func main() {
    var a, b, s int
    fmt.Scan(&a, &b)
    if a > b {
        fmt.Println("invalido")
        return
    }
    for ; a <= b; a ++ {
        if a%2 == 0 && a%3 == 0 {
            s+= a 
        }
    }
    fmt.Println(s)
}
