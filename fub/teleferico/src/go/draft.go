package main
import "fmt"
func main() {
    var C, A int
    fmt.Scan(&C, &A)

    if (A%(C-1) == 0) {
        fmt.Println(A/(C-1))
    } else {
        fmt.Println(A/(C-1) + 1)
    }


}
