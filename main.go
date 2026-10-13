package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: portman <port>")
		return
	}

	fmt.Println(os.Args[1])
}
