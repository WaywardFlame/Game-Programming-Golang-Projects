package main

import (
	"fmt"
	"math/rand/v2"
)

type creature struct {
	name         string
	dexterity    int
	strength     int
	intelligence int
}

func main() {
	creatureOne := NewRPGCreature("Alkhait")
	creatureTwo := RandomRPGCreature("Logance")
	fmt.Println(creatureOne)
	fmt.Println(creatureTwo)
}

func NewRPGCreature(name string) creature {
	return creature{name: name, dexterity: 1, strength: 1, intelligence: 1}
}

func RandomRPGCreature(name string) creature {
	return creature{name: name, dexterity: rand.IntN(9) + 1, strength: rand.IntN(9) + 1, intelligence: rand.IntN(9) + 1}
}
