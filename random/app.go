package main

import (
	"random/week2"
)

func main() {
	go week2.RoutineTest("Hello, Goroutine!")
	week2.RoutineTest("Hello, Main!")
}
