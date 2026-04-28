package main
import "fmt"
// import "math/rand"

// 0 -pedra
// 1 papel
// 2 tesoura

func main() {
    // valor := rand.Intn(5) // serve para jogar um número aleatório
    // fmt.Println(valor)

    const PEDRA int = 0
    const PAPEL int = 1
    const TESOURA int = 2

    var jog1, jog2 int
    fmt.Println("Jog 1: Digite 0 (pedra), 1 (papel), 2 (tesoura)")
    fmt.Scan(&jog1)
    fmt.Println("Jog 2 Digite 0 (pedra), 1 (papel), 2 (tesoura)")
    fmt.Scan(&jog2)

    fmt.Println(jog1, jog2)

}
