package main

import "fmt"

func Struct() {
	sandra := newFemale("Sandra", 33)
	otto := newMale("Otto", 40)
	sandra.Print()
	fmt.Println("-----------")
	otto.Print()
}

func newMale(name string, age int) *Person {
	return &Person{name: name, age: age, sex: "male"}
}

func newFemale(name string, age int) *Person {
	return &Person{name: name, age: age, sex: "female"}
}

type Person struct {
	name string
	age  int
	sex  string
}

func (p *Person) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("sex: ", p.sex)
}
