 package main
import "fmt"

func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    diferenca := n1 - n2
    
    reseultado := abs(diferenca)

    fmt.Println(reseultado)


}
