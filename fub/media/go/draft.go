package main
import "fmt"
func main() {
    var a, b float64
    var media float64
    fmt.Scan(&a, &b)
    media = (a + b) / 2 // serve para fazer a média
    fmt.Printf("%.1f\n", media) // serve para deixar o número como decimal
}
