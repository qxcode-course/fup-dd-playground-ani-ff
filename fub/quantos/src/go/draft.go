package main
import "fmt"
func main() {
    var j1, j2, j3 int
    fmt.Scan(&j1, &j2, &j3)

    if j1 == j2 && j2 == j3 {
        fmt.Println(3)
    } else if j1 == j2 || j1 == j3 || j2 == j3 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}
