package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, World!", runtime.GOOS)
}
