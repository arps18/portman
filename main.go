package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: portman <port>")
		return
	}

	port, err := strconv.Atoi(os.Args[1])
	// fmt.Println(os.Args[1])
	if err != nil {
		fmt.Println("Incorrect port number")
		return
	}

	fmt.Println(port)
}
