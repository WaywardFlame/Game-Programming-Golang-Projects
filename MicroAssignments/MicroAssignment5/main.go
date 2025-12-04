package main

import (
	"fmt"
	"math/rand/v2"
)

type Creature struct {
	name         string
	dexterity    int
	strength     int
	intelligence int
}

func main() {
	randomSlice := CreateRPGRoster(5)
	fmt.Println(randomSlice)
}

func RandomRPGCreature(name string) Creature {
	return Creature{name: name, dexterity: rand.IntN(9) + 1, strength: rand.IntN(9) + 1, intelligence: rand.IntN(9) + 1}
}

func CreateRPGRoster(size int) []Creature {
	slice := make([]Creature, size)
	var dex int = 0
	var str int = 0
	var intel int = 0
	for i := 0; i < size; i++ {
		slice[i] = RandomRPGCreature("RandomCreature")
		dex += slice[i].dexterity
		str += slice[i].strength
		intel += slice[i].intelligence
	}
	fmt.Println("Total Dexterity:", dex)
	fmt.Println("Total Strength:", str)
	fmt.Println("Total Intelligence:", intel)
	return slice
}
