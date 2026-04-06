package main
import "fmt"
func main() {
    const PEDRA int = 0
    const PAPEL int = 1
    const TESOURA int = 2

    var jog1, jog2 int
    fmt.Println("Jog1: Digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog1)
    fmt.Println("Jog2: Digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog2)

    if jog1 == jog2 {
        fmt.Println("Empate")
    } else if (jog1 == PEDRA && jog2 == TESOURA) ||
        (jog1 == PAPEL && jog2 == PEDRA) ||
        (jog1 == TESOURA && jog2 == PAPEL) {
        fmt.Println("Jog1 ganhou")
    } else {
        fmt.Println("Jog2 ganhou")
    }
}
