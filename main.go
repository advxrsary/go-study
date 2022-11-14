// Golang basics tutorial

package main

import (
	"fmt"
	"os"
)

// func to decrement two numbers from command line arguments
func decrement(a []string, b []string) int {
	var x, y int
	fmt.Sscanf(a[0], "%d", &x)
	fmt.Sscanf(b[0], "%d", &y)
	return x - y
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: main.go <number1> <number2>")
		return
	}
	fmt.Println(decrement(os.Args[1:], os.Args[2:]))
}
