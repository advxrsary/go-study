package hash

// Given a personal information such as an email address, IP address, or a public key, the program you will write needs to generate a unique avatar.
// Imagine that you are building a new application and you want all of your users to have a default and unique avatar.
// The package you will write will allow the generation of such avatars.
// GitHub recently used such an approach and generates an identicon for all new users who don't have a gravatar account attached.

// This program generates a random avatar using user's IP address declared in the environment variable IP.
// The program uses the hash package to generate a hash from the IP address.
// The hash is then used to generate random shapes and colors. The program uses the image package to create a new image and draw the shapes on it.
// The image is then saved as avatar.png file.

import (
	"crypto/sha256"
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/andelf/go-curl"
)

func CurlGetIp() string {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://ifconfig.me/ip")
	easy.Setopt(curl.OPT_FOLLOWLOCATION, 1)

	var ip string
	easy.Setopt(curl.OPT_WRITEDATA, &ip)

	easy.Perform()
	return ip
}

func GenerateAvatar(ip string) {
	// create a new image
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	// generate background color
	hash := sha256.Sum256([]byte(ip))
	rand.Seed(int64(hash[0]))
	bgColor := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			img.Set(i, j, bgColor)
		}
	}

	// generate random seed
	hasher := sha256.New()
	hasher.Write([]byte(ip))
	seed, _ := hex.DecodeString(hex.EncodeToString(hasher.Sum(nil)))
	rand.Seed(int64(seed[0]))

	// draw random shapes
	for i := 0; i < 10; i++ {
		x := rand.Intn(300)
		y := rand.Intn(300)
		width := rand.Intn(300)
		height := rand.Intn(300)
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		a := 255
		shape := rand.Intn(2)
		if shape == 0 {
			drawCircle(img, x, y, width, height, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		} else {
			drawRectangle(img, x, y, width, height, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	// draw random squares
	for i := 0; i < 10; i++ {
		x := rand.Intn(300)
		y := rand.Intn(300)
		width := rand.Intn(300)
		height := rand.Intn(300)
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		a := 255
		drawSquare(img, x, y, width, height, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
	}

	// draw random triangles
	for i := 0; i < 10; i++ {
		x := rand.Intn(300)
		y := rand.Intn(300)
		width := rand.Intn(300)
		height := rand.Intn(300)
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		a := 255
		drawTriangle(img, x, y, width, height, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
	}

	// save the image
	f, _ := os.Create("avatar.png")
	png.Encode(f, img)
}

func drawCircle(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if (i-x)*(i-x)+(j-y)*(j-y) <= width*width/4 {
				img.Set(i, j, c)
			}
		}
	}
}

func drawRectangle(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i >= x && i <= x+width && j >= y && j <= y+height {
				img.Set(i, j, c)
			}
		}
	}
}

func drawSquare(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i >= x && i <= x+width && j >= y && j <= y+width {
				img.Set(i, j, c)
			}
		}
	}
}

func drawTriangle(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i >= x && i <= x+width && j >= y && j <= y+width {
				img.Set(i, j, c)
			}
		}
	}
}
