package main

import "fmt"

type MagicAnimal interface {
	Speak() string
}

type Rabbit struct {}

func (r Rabbit) Speak() string {
	return "How long is forever?"
}

func main() {
	animals := []MagicAnimal{Rabbit{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
