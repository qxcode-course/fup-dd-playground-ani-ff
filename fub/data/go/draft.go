package main
import "fmt"
func main() {
    var hr, min int
    fmt.Scan(&hr, &min)

    var dia, mes, ano int
    fmt.Scan(&dia, &mes, &ano)

    // serve para adicionar os : e colocar o 0 antes do número int
    fmt.Printf("%02d:%02d %02d/%02d/%002d\n", hr, min, dia, mes, ano%100)

}
