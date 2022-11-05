package main

import "fmt"

type Sayer interface {
	SaySomething()
}

type Animal struct {
	Weight int
}

func (a Animal) SaySomething() {
	fmt.Println("йоу!")
}

type Dog struct {
	Animal
}

func (d Dog) SaySomething() {
	fmt.Println("гав!")
}

type Cat struct {
	Animal
}

func (c Cat) SaySomething() {
	fmt.Println("мяу!")
}

type MagicAnimal struct {
	Animal
}

func main() {
	var a Sayer = Animal{Weight: 10}
	var d Sayer = Dog{Animal{Weight: 10}}
	var c Sayer = Cat{Animal{Weight: 10}}

	a.SaySomething()
	d.SaySomething()
	c.SaySomething()

	var m Sayer = MagicAnimal{Animal{10}}
	m.SaySomething()

	var c1 Sayer = c.(Animal) // not working
	c1.SaySomething()
}
