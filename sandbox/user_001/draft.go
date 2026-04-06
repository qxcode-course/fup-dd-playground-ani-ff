package main
import "fmt"


func diga_oi() string {
    return "oi"
}

// sempre vai começar o main primeiro
func main() {
    fmt.Println("a")
    valor := diga_oi()
    diga_oi()
    fmt.Println("b")
}