// Golang basics tutorial

package main

import (
	"fmt"
	"os"
	"strings"
)

// efficent way to do it
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
