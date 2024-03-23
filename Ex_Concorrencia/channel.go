package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second)
	}
}

func printar(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go printar(c)

	var escreva string
	fmt.Scanln(&escreva)
}
