package main

import (
	"go-study/hash"
	"go-study/loading_anim"
)

func main() {
	loading_anim.LoadAnim(200)
	hash.GenerateAvatar()
}
