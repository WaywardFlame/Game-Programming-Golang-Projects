package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//I considered seeding rand, but the tests on
	//my machine led to different results. Looking
	//into it, it seems that Go may have added in
	//the auto-seeding of rand at some point.
	//Additionally, rand v2 does not use the typical
	//Seed() function. Thus, I chose to avoid attempting
	//to seed rand for this micro assigment.

	fmt.Println("Welcome to Simon Says.")
	fmt.Println("If Simon says a number, enter that number.")
	fmt.Println("Else, enter a different number.")
	fmt.Println()

	var randomNum int
	var simonSays int
	var playerNum int
	for {
		randomNum = rand.IntN(1000) // generate random num from 0 to 999
		simonSays = rand.IntN(2)    // generates 0 or 1
		if simonSays == 1 {
			fmt.Println("simon says", randomNum)
		} else {
			fmt.Println(randomNum)
		}

		fmt.Scan(&playerNum)
		if simonSays == 1 && playerNum == randomNum {
			continue
		} else if simonSays == 1 && playerNum != randomNum {
			fmt.Println("Game Over")
			break
		} else if simonSays == 0 && playerNum != randomNum {
			continue
		} else {
			fmt.Println("Game Over")
			break
		}
	}
}
