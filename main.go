// Golang basics tutorial

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("\nvars:", s)
}

func echo2() {
	fmt.Println("\nvars: ")
	for i, arg := range os.Args[1:] {
		fmt.Printf("#%d\t-\t%s\n", i, arg)
	}
}

func echo3() {
	fmt.Printf("\nvars: %s", strings.Join(os.Args[1:], " "))
}

// efficent way to do it
func main() {
	fmt.Println("name: ", os.Args[0])

	start := time.Now()
	echo1()
	secs := time.Since(start).Seconds()
	fmt.Printf("took: %fs\n", secs)

	start = time.Now()
	echo2()
	secs = time.Since(start).Seconds()
	fmt.Printf("took: %fs\n", secs)

	start = time.Now()
	echo3()
	secs = time.Since(start).Seconds()
	fmt.Printf("\ntook: %fs\n", secs)

}
