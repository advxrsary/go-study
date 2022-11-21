package loading_anim

import (
	"fmt"
	"time"
)

// func to display three dots in sqauare braces loading animation
// function should stop when hide_loading() is called

// Python version:
// print("Loading:")

// #animation = ["10%", "20%", "30%", "40%", "50%", "60%", "70%", "80%", "90%", "100%"]
// animation = ["[■□□□□□□□□□]","[■■□□□□□□□□]", "[■■■□□□□□□□]", "[■■■■□□□□□□]", "[■■■■■□□□□□]", "[■■■■■■□□□□]", "[■■■■■■■□□□]", "[■■■■■■■■□□]", "[■■■■■■■■■□]", "[■■■■■■■■■■]"]

// for i in range(len(animation)):
//     time.sleep(0.2)
//     sys.stdout.write("\r" + animation[i % len(animation)])
//     sys.stdout.flush()

// Go version:
// print("\n")

// - Hey, Copilot! How do I remove % sign from console output?
// + You need to use fmt.Printf() function to print to console. in scientific notation.

func Animation(duration int) {
	animation := []string{"[■□□□□□□□□□]", "[■■□□□□□□□□]", "[■■■□□□□□□□]", "[■■■■□□□□□□]", "[■■■■■□□□□□]", "[■■■■■■□□□□]", "[■■■■■■■□□□]", "[■■■■■■■■□□]", "[■■■■■■■■■□]", "[■■■■■■■■■■]"}
	for i := 0; i < len(animation); i++ {
		time.Sleep(time.Duration(duration) * time.Millisecond)
		fmt.Printf("\r%s", animation[i%len(animation)])
	}
}

func HideLoading() {
	fmt.Printf("\r\x1b[K")
}

func LoadAnim(duration int) {
	Animation(duration)
	HideLoading()
	fmt.Println("./avatar.png generated")
}
