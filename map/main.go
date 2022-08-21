package main

import "fmt"

func main() {
	colors := map[string]string{ // Declare a map
		"red":   "#ff0000",
		"green": "#4b5f5r",
		"white": "#ffffff",
	}

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}