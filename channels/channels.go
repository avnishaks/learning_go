package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func Channels() {
	fmt.Println("Welcome to world of channels")
	/*

		Introduction about channels :-
		1. Hold Data
		2. Channels is thread safe
		3. Channels listen for data

	*/

	//var c = make(chan int)
	//c <- 1
	//var i = <-c
	//fmt.Println(i)
	// above code result in the deadlock situation , we should use channel with the help of goroutine
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
	chickenChannels()

}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

var MAX_CHICKEN_PRICE float32 = 5

func chickenChannels() {
	var chickenchannel = make(chan string)
	var website = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range website {
		go chickenPrice(website[i], chickenchannel)
	}
	sendMessage(chickenchannel)
}

func chickenPrice(website string, chickenchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrices = rand.Float32() * 20
		if chickenPrices <= MAX_CHICKEN_PRICE {
			chickenchannel <- website
			break
		}
	}
}

func sendMessage(chickenchannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenchannel)
}
