package main
import "fmt"
func main() {
    var n, d, a int
    fmt.Scan(&n, &d, &a)
    var resultado int 
    if d >= a {
        resultado = d - a
    } else {
        resultado = n - a + d
    }
    fmt.Println(resultado)
}
