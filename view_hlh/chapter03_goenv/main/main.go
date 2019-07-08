package main

import (
	"fmt"
	"os"
)

func main() {
	var gos = os.Getenv("GOOS")
	fmt.Printf("The operating os is : %s\n", gos)
	path := os.Getenv("GOPATH")
	fmt.Printf("The Gopath is : %s\n", path)
}
