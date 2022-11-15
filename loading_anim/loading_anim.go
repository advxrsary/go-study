package loading_anim

import (
	"fmt"
	"time"
)

// func to display three dots in sqauare braces loading animation
// function should stop when hide_loading() is called
func ShowLoading() {
	for {
		fmt.Print("\r[   ]")
		time.Sleep(time.Millisecond * 250)
		fmt.Print("\r[.  ]")
		time.Sleep(time.Millisecond * 250)
		fmt.Print("\r[.. ]")
		time.Sleep(time.Millisecond * 250)
		fmt.Print("\r[...]")
		time.Sleep(time.Millisecond * 250)

	}
}

// func to stop the loading animation, and clear the line
func HideLoading(start time.Time) {

	fmt.Printf("\r[%s]", time.Since(start))
}

func LoadAnim(milliseconds time.Duration) {
	start := time.Now()
	go ShowLoading()
	time.Sleep(time.Millisecond * milliseconds)
	HideLoading(start)
}
