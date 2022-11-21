package main

import (
	"go-study/hash"
	"go-study/loading_anim"
)

func main() {
	ip := hash.CurlGetIp()
	loading_anim.LoadAnim(40)
	hash.GenerateAvatar(ip)
}
