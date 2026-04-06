package main

import (
	"fmt"
)
func main() {
    const pedra int = 0
    const PAPEL int = 1
    const TESOURA int = 2

    var jog1, jog2 int
    fmt.Println("Jog1: Digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog1)
    fmt.Println("Jog2: Digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog2)


}
