package main
import "fmt"
import "math"

func round(x float64) int {
    return int(math.Round(x))
}

func floor(x float64) int {
    return int(math.Floor(x))
}

func ceil(x float64) int {
    return int(math.Ceil(x))
}

func main() {
    var operacao string
    var num float64
    fmt.Scan(&operacao)
    fmt.Scan(&num)

    switch {
    case operacao == "r":
        fmt.Println(round(num))
    case operacao == "f":
        fmt.Println(floor(num))
    case operacao == "c":
        fmt.Println(ceil(num))
    }
}
