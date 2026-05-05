package main
import "fmt"
func main() {
    var A, B, C int
    fmt.Scan(&A, &B, &C)

    delta := B * B - 4 * A *C

    if delta {
        fmt.Println
    }

    fmt.Println(delta)
}
