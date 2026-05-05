package main
import "fmt"
func main() {
    var A, B, C, H, L int
    fmt.Scan(&A, &B, &C, &H, &L)

    if A <= H && A <= L {
        fmt.Println("S")
    } else if B <= H && B <= L {
        fmt.Println("S")
    } else if C <= H && C <= L {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }

}
