package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')

	for i := len(s) - 2; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
    fmt.Println("")
}