package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    divisao := n1 / n2
    resto := n1%n2
    divisao2 := float64(n1) / float64(n2) // mesmo resultado da divisão normal, mas com os 00 dps

    
    fmt.Println(divisao)
    fmt.Println(resto)
    fmt.Printf("%.2f\n", divisao2)
}
