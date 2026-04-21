package main
import "fmt"
func main() {
    var P, D1, D2 int
    fmt.Scan(&P, &D1, &D2)

    soma := D1 + D2

    if soma % 2 == 0 && P == 0 {
        fmt.Println(0)
    } else if soma % 2 == 0 && P == 1 {
        fmt.Println(1)
    } else if soma % 2 != 0 && P == 0 {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }

}
