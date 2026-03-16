package main
import "fmt"
func main() {
    var a, b float64
    var media float64
    fmt.Scan(&a, &b)
    media = (a + b) / 2
    fmt.Printf("%.1f\n", media)
}
