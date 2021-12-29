package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(msg string) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println(msg + " - goroutine")
}

func main() {
	go hello("Hello 1")
	go hello("Hello 2")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("chamada normal")

}
