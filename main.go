package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, World")
}

func nowTime() {
	fmt.Println("Welcome to the playground")
	fmt.Println("The time is", time.Now())
}

func main() {
	hello()
	nowTime()
}
