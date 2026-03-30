package main
import "fmt"
func main() {
    var tempo int
    var hora int
    var resto int
    var minutos int
    var resto2 int
    var segundos int
    
    fmt.Scan(&tempo, &hora, &resto, &minutos, &resto2, &segundos)

    hora = tempo / 3600
    resto = tempo % 3600
    minutos = resto / 60
    resto2 = resto % 60
    segundos = resto2

    fmt.Printf("%d:%d:%d\n", hora, minutos, segundos)


}
