package main

import "fmt"

const CatShout = "Meow"
const DogShout = "Wuff"

func Interface() {
	cat := Cat{}
	fmt.Print("Cats shout: ")
	cat.shout()
	fmt.Println("...but")
	dog := Dog{}
	fmt.Print("Dogs shout: ")
	dog.shout()
}

type Dog struct {
}

func (d *Dog) shout() {
	fmt.Println(DogShout)
}

type Cat struct {
}

func (c *Cat) shout() {
	fmt.Println(CatShout)
}

type Animal interface {
	shout() string
}
