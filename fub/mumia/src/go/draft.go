package main
import "fmt"
func main() {
    var nome string
    var idade int
    fmt.Scan(&nome, &idade)

    crianca := idade < 12

    if idade < 12 {
        fmt.Println(nome, "eh", crianca)
    }
}
