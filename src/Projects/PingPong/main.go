package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for {
		c <- "ping"
	}
}
func pongar(c chan string) {
	for {
		c <- "pong"
	}
}

func printar(c chan string) {
	for {
		msg := <-c

		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go printar(c)
	go pongar(c)

	var escreva string
	fmt.Scanln(&escreva)
}
