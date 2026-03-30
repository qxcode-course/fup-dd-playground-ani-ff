package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    soma := a + b
    subtracao := a - b
    multiplicacao := a * b
    divisao := float64(a) / float64(b)
    resto := a % b

    fmt.Println(soma)
    fmt.Println(subtracao)
    fmt.Println(multiplicacao)
    fmt.Printf("%.2f\n", divisao)
    fmt.Println(resto)
}
