package main
import "fmt"
func main() {
    var celsius float64
    var fah float64

    fmt.Scan(&celsius, &fah)

    fah = 1.8 * celsius + 32

    fmt.Printf("%.6f\n", fah)
}
