package main
import "fmt"
func main() {
    var c byte
    fmt.Scanf("%c", &c)

    if c >= 'a' && c <= 'z' {
        c -= 32
    } else if c >= 'A' && c <= 'Z' {
        c += 32
    }
    fmt.Printf("%c\n", c)
}
