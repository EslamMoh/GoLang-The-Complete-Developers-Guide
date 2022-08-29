package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // This is how we create new channel. make is a built-in function that will create a value out of the given type.

	for _, link := range links {
		go checkLink(link, c) // We pass the channel to the Go Routine definitions that will use it
	}

	// for i := 0; i < len(links); i++ { // This is how we will keep the main thread busy by print the the value from the channel, through iterating by the count of the slice
	// 	fmt.Println(<-c)
	// }

	for l := range c { // Wait for the channel to return some value and after the channel has returned some value, assign it to the variable l, then run the body inside the for loop to make it always spawn new Go routine
		// This can be done through inifinit loop but this format for the loop is more understandble.
		// go checkLink(l, c) // We pass c channel again
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // these are the paranthese that actually execute literal function
	}
}

func checkLink(link string, c chan string) { // We define that the method will recive c var with type channel and the channel of type string.
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // We pushed the link again to the channel to make sure it will kick off new Go routine using this link, to be run repeatedly
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
