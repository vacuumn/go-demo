package main

import (
	"fmt"
	"../example"
)


func main() {
	animals := []example.MagicAnimal{
			example.Rabbit{"White rabbit", 0},
		}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}


