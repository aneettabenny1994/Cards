package main

import "fmt"

func newMap() {

	var colorsMap map[string]string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	morecolors := make(map[int]string)

	morecolors[10] = "#ffffff"

	delete(morecolors, 10)

	fmt.Println(colorsMap)
	printMap(colors)
	fmt.Println(morecolors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
