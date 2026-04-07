package main
import "fmt"
func main() {
    var n1, n2 int
    var c string
    fmt.Scan(&n1, &n2, &c)

    resultado := 0

    // if c == "+" {
    //     resultado = n1 + n2
    //     fmt.Println(resultado)
    // } else if c == "-" {
    //     resultado = n1 - n2
    //     fmt.Println(resultado)
    // } else if c == "*" {
    //     resultado = n1 * n2
    //     fmt.Println(resultado)
    // } else if c == "/" {
    //     resultado = n1 / n2
    //     fmt.Println(resultado)
    // }

    switch {
    case c == "+":
        resultado = n1 + n2
        fmt.Println(resultado)
    case c == "-":
        resultado = n1 - n2
        fmt.Println(resultado)
    case c == "*":
        resultado = n1 * n2
        fmt.Println(resultado)
    case c == "/":
        resultado = n1 / n2
        fmt.Println(resultado)
    }
}