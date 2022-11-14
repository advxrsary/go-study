// Golang basics tutorial

package main

import (
	"fmt"
	"os"
)

// non-efficent way to do it
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
