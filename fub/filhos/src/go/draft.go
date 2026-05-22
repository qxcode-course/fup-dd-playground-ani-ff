package main
import "fmt"
func main() {
    var idade, filhos int
    fmt.Scan(&idade, &filhos)
    for i := 0; i < filhos; i++ {
        fmt.Println(idade + i*2)
    }
}
