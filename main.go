package main

import (
	"go-study/echo"
	"go-study/loading_anim"
)

func main() {
	// this function takes duration of the animation in milliseconds as an argument
	loading_anim.LoadAnim(2000)
	echo.Echo2()

}
