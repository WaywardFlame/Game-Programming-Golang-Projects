package main

import (
	"fmt"
)

func main() {
	var myLevel int = 1
	var xp int = 34
	myLevel, xp = AwardXP(myLevel, xp, 174)
	fmt.Println("My current level is:", myLevel)
	fmt.Println("My current xp is:", xp)
}

func AwardXP(currentLevel int, currentXP int, awardedXP int) (int, int) {
	if awardedXP < 0 {
		fmt.Println("Warning: Can't award positive XP. Returning original values.")
		return currentLevel, currentXP
	}

	currentXP += awardedXP
	currentLevel += currentXP / 100
	currentXP %= 100
	return currentLevel, currentXP
}
