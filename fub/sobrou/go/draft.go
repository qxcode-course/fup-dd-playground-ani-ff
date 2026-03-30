package main
import "fmt"
func main() {
    var q1, q2, q3 int
    var p1, p2, p3 float64
    var dinheiro float64

    fmt.Scan(&q1, &q2, &q3) // Entrada das quantidades
    fmt.Scan(&p1,&p2, &p3) // Entrada dos preços
    fmt.Scan(&dinheiro) // Entrada do dinheiro

    ctotal := float64(q1)*p1 + float64(q2)*p2 + float64(q3)*p3 // Calculo do custo total

    troco := dinheiro - ctotal // Calculo do troco

    fmt.Printf("%.2f\n", troco)

    
}
