package main
import "fmt"
func main() {
    var velocidade float64
    var tempo float64
    var combustivel float64

    fmt.Scan(&velocidade, &tempo, &combustivel)

    tempoh := tempo / 60
    distancia := velocidade * tempoh
    desempenho := distancia / combustivel

    fmt.Printf("%.2f\n", desempenho)
}
