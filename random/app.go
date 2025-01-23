package main

import (
	"fmt"
	"strconv"
	"time"
)

func WarmUp() {
	// Warmup
	fmt.Println("Hello, World!")

	var something string = "big something";

	aNumb := 10
	
	var final = something + strconv.Itoa(aNumb)
	// var final = something + string(aNumb)
	fmt.Println(final)
}

func main(){
	// go RoutineTest("Hello, Goroutine!")
    // RoutineTest("Hello, Main!")
}

func RoutineTest(str string) {
	time.Sleep(1 * time.Second)
	fmt.Println(str)
}

func PlayingWithPointers(){
	num := 1
	fmt.Println(&num)
}

func PlayingWithObjects(){
	// var p Person = Person{name: "John", age: 25}
	// var p *Person = new(Person)

	// Reference
	var p Person = *new(Person)

	fmt.Println(&p)
	p2 := p

	// Unreference
	// p2 := *p
	fmt.Println(&p2)

	p2.name = "John"
	fmt.Println(p.greet())
	fmt.Println(p2.greet())
}

type Person struct {
	name string
	age int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.name + " and I am " + strconv.Itoa(p.age) + " years old."
}