package main
import "fmt"
func main() {
    var C, banana, goiaba, manga int
    fmt.Scan(&C, &banana, &goiaba, &manga)

    total := banana + goiaba + manga
    min := total / C

    if  total % C != 0{
        fmt.Println(min+1)
    } else {
        fmt.Println(min)
    }
}
