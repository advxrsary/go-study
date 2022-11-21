package main

import (
	"container/list"
	"fmt"
	"go-study/loading_anim"
	"os"
)

type Vehicle struct {
	Brand, Model, Color string
}

// create a list of brands, models and colors

var (
	brands = []string{"BMW", "Audi", "Mercedes", "Volkswagen", "Porsche", "Ferrari", "Lamborghini", "Maserati", "Bugatti", "Tesla"}
	colors = []string{"Black", "White", "Red", "Blue", "Green", "Yellow", "Orange", "Purple", "Pink", "Brown"}
	models = []string{"X5", "A6", "C63", "Golf", "911", "488", "Huracan", "Ghibli", "Chiron", "Model S"}
)

// create a function that will take os.Args in random order as input and match it with the list of brands, models and colors and return a Vehicle struct with the matched values
// if there is no match, return an error saying that the car was not found
func MatchArgs(args []string) Vehicle {
	var car Vehicle
	var l = list.New()
	for _, arg := range args {
		l.PushBack(arg)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		for _, brand := range brands {
			if e.Value == brand {
				car.Brand = e.Value.(string)
			}
		}
		for _, model := range models {
			if e.Value == model {
				car.Model = e.Value.(string)
			}
		}
		for _, color := range colors {
			if e.Value == color {
				car.Color = e.Value.(string)
			}
		}
	}
	return car
}

func main() {
	// print vehicle struct
	fmt.Print("\t Looking for your car")

	loading_anim.LoadAnim(5134)
	if len(os.Args) == 1 {
		fmt.Println("Please provide a brand, model and color")
		return
	}
	fmt.Printf("\nYour car is a %s %s %s\n", MatchArgs(os.Args).Color, MatchArgs(os.Args).Brand, MatchArgs(os.Args).Model)
}
