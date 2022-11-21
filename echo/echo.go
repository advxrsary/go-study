package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() []string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return strings.Split(s, " ")
}

func Echo2() {
	fmt.Println("\nvars: ")
	for i, arg := range os.Args[1:] {
		fmt.Printf("#%d\t-\t%s\n", i, arg)
	}
}

func Echo3() {
	fmt.Printf("\nvars: %s", strings.Join(os.Args[1:], " "))
}
