package main
import "fmt"
func main() {
    var nome string
    var idade int

    for {
        fmt.Println("Digite nome e idade")
        fmt.Scan(&nome, &idade)
        fmt.Println("Seu nome é %v e idade %v\n", nome)
    }

}