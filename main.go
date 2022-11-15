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
	fmt.Println(os.Args[0], s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(os.Args[0], s)
}

func echo3() {
	fmt.Println(os.Args[0], strings.Join(os.Args[1:], " "))
}

// efficent way to do it
func main() {
	start := time.Now()
	echo1()
	secs := time.Since(start).Seconds()
	fmt.Println(secs)

	start = time.Now()
	echo2()
	secs = time.Since(start).Seconds()
	fmt.Println(secs)

	start = time.Now()
	echo3()
	secs = time.Since(start).Seconds()
	fmt.Println(secs)

}
