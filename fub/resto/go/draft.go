package main
import "fmt"
func main() {
    var bombons, criancas int
    fmt.Scan(&bombons, &criancas)
    fmt.Println(bombons/criancas, bombons%criancas) // serve para fazer a divisão de 2 números e o resto
}
