package main

import "fmt"

type person struct {
	name string
	age int
}

func send(ch chan<- person, p person) {
	ch <- p
}

func main() {
	c := make(chan person)

	go send(c, person{name: "John", age: 31} )

	a  := <- c
	fmt.Println(a.name)
	fmt.Println(a.age)
	close(c)
}
