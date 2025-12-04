package main

import (
	"fmt"
	"math/rand/v2"
)

type character struct {
	fortressHealth      int
	items               [3]int
	tickDamage          int
	isSmoked            bool
	isDefending         bool
	successiveVictories int
	playerChoice        string
	aiChoice            int
	aiRPSDecision       string
	aiItemDecision      string
	aiRPSWeight         int
}

// I could have had one decision variable, but
// I couldn't be bothered to change it by the
// time I realized that.

func main() {
	startMenu()
	fortressMode()
}

func startMenu() {
	fmt.Println("Welcome to the Rock-Paper-Scissors Fortress Game")
	fmt.Println("-------------------------------------------")
	fmt.Println("Type 'rock', 'paper', 'scissors', 'defend', or 'items' to choose.")
	fmt.Println("After 'items', type 'fire_bomb', 'cannon_ball', or 'smoke_screen' to choose.")
	fmt.Println("Type 'manual' to bring up information (when not choosing item).")
	fmt.Println("Type 'exit' to end the game.")
	fmt.Println("Begin!")
	fmt.Println("-------------------------------------------")
}

func manual() {
	fmt.Println("-------------------------------------------")
	fmt.Println("Rock-Paper-Scissors Fortress:")
	fmt.Println("A RPS game in which you and the AI will attempt to defend a")
	fmt.Println("fortress from each other. Both fortresses have a limited amount")
	fmt.Println("of health. You will have several options available for you to")
	fmt.Println("use for either defending your fortress, or destroying the enemy's")
	fmt.Println("fortress.")
	fmt.Println("    - RPS: You have the standard rock-paper-scissors choice, with")
	fmt.Println("           draws dealing a small amount of damage to both sides. Successive")
	fmt.Println("           victories at RPS will deal more and more damage to the other side.")
	fmt.Println("    - Defend: By typing 'defend' to defend, you can negate incoming RPS damage.")
	fmt.Println("           This does not end successive victory streaks, nor does it stop item")
	fmt.Println("           damage.")
	fmt.Println("    - Items: By typing 'items', you can choose which items to use against the")
	fmt.Println("           enemy. Items bypass enemy defenses, and are awarded randomly to one")
	fmt.Println("           side at the start of each round, except for the first round.")
	fmt.Println("    - Fire Bomb (item): By typing 'fire_bomb' after typing in 'items', you will")
	fmt.Println("           use a fire bomb. Fire bombs deal 5 tick damage for 5 rounds, starting")
	fmt.Println("           in the round it is first used. Using a fire bomb when one is already")
	fmt.Println("           active will only reset the tick damage counter.")
	fmt.Println("    - Cannon Ball (item): By typing 'cannon_ball' after typing in 'items', you")
	fmt.Println("           will use a cannon ball. Cannon balls deal 20 damage.")
	fmt.Println("    - Smoke Screen (item): By typing 'smoke_screen' after typing in 'items, you")
	fmt.Println("           will use a smoke screen. Smoke screens deal no damage, but prevent the")
	fmt.Println("           the opposing side from using RPS for one round.")
	fmt.Println("    - Successive Victories: Repeated victories at RPS will increase the damage")
	fmt.Println("           that is dealt to the opposing side. Each victory add 5 additional")
	fmt.Println("           damage, and caps out at 15 additional damage (3 victories). Losing")
	fmt.Println("           RPS will reset the victories and additional damage.")
	fmt.Println("-------------------------------------------")
}

func fortressMode() {
	// struct to determine player status
	player := character{fortressHealth: 100, tickDamage: 0, isSmoked: false, isDefending: false, successiveVictories: 0}

	// struct to determine ai status
	// ac = ai character
	ai := character{fortressHealth: 100, tickDamage: 0, isSmoked: false, isDefending: false, successiveVictories: 0, aiRPSWeight: 0}

	// variable to use in loop
	var i int = 1

	// the game loop
	for ; ; i++ {
		fmt.Println("")
		fmt.Println("Round", i)
		fmt.Println("")

		printStatus(0, player)
		printStatus(1, ai)

		fmt.Println("Your Move:")
		playerMove(&player, &ai)
		if player.playerChoice == "exit" {
			break
		}

		fmt.Println("AI Move:")
		aiMove(&ai)
		if ai.aiChoice < 5 {
			fmt.Println(ai.aiItemDecision)
		} else {
			fmt.Println(ai.aiRPSDecision)
		}

		// adjust health, determine victory based on selected moves
		// objected oriented programming would probably make this easier
		determineResults(&player, &ai)

		// adjust status of player and AI
		adjustStatus(&player, &ai)

		// check for player win or ai win
		if player.fortressHealth <= 0 {
			fmt.Println("")
			fmt.Println("The AI wins!")
			break
		} else if ai.fortressHealth <= 0 {
			fmt.Println("")
			fmt.Println("The Player wins!")
			break
		}
	}
}

func checkRPS(playerChoice string, aiDecision string) int {
	// 0 - draw, 1 - player lose ai win, 2 - player win ai lose
	if playerChoice == "rock" {
		if aiDecision == "rock" { // rock
			fmt.Println("RPS: DRAW")
			return 0
		} else if aiDecision == "paper" { // paper
			fmt.Println("RPS: PLAYER LOSE, AI WIN")
			return 1
		} else if aiDecision == "scissors" { // scissors
			fmt.Println("RPS: PLAYER WIN, AI LOSE")
			return 2
		}
	} else if playerChoice == "paper" {
		if aiDecision == "rock" { // rock
			fmt.Println("RPS: PLAYER WIN, AI LOSE")
			return 2
		} else if aiDecision == "paper" { // paper
			fmt.Println("RPS: DRAW")
			return 0
		} else if aiDecision == "scissors" { // scissors
			fmt.Println("RPS: PLAYER LOSE, AI WIN")
			return 1
		}
	} else if playerChoice == "scissors" {
		if aiDecision == "rock" { // rock
			fmt.Println("RPS: PLAYER LOSE, AI WIN")
			return 1
		} else if aiDecision == "paper" { // paper
			fmt.Println("RPS: PLAYER WIN, AI LOSE")
			return 2
		} else if aiDecision == "scissors" { // scissors
			fmt.Println("RPS: DRAW")
			return 0
		}
	} else {
		fmt.Println("Error encountered")
	}
	return -1
}

func determineResults(player *character, ai *character) {
	// check RPS
	// if both chose RPS
	if (player.playerChoice == "rock" || player.playerChoice == "paper" || player.playerChoice == "scissors") && ai.aiChoice >= 5 {
		rpsResult := checkRPS(player.playerChoice, ai.aiRPSDecision)
		if rpsResult == 0 { // draw
			player.fortressHealth -= 5
			ai.fortressHealth -= 5
			player.successiveVictories = 0
			ai.successiveVictories = 0
		} else if rpsResult == 1 { // player lose, ai win
			player.fortressHealth -= (10 + 5*ai.successiveVictories)
			player.successiveVictories = 0
			ai.successiveVictories += 1
			ai.aiRPSWeight += 1
			if ai.successiveVictories > 3 { // limit ai successive victories
				ai.successiveVictories = 3
			}
			if ai.aiRPSWeight > 3 { // limit ai RPS weight
				ai.aiRPSWeight = 3
			}
		} else if rpsResult == 2 { // player win, ai lose
			ai.fortressHealth -= (10 + 5*player.successiveVictories)
			ai.successiveVictories = 0
			ai.aiRPSWeight -= 1
			player.successiveVictories += 1
			if player.successiveVictories > 3 { // limit player successive victories
				player.successiveVictories = 3
			}
			if ai.aiRPSWeight < 0 { // limit ai RPS weight
				ai.aiRPSWeight = 0
			}
		}
		// if player chose RPS, and ai did not choose RPS or defend
	} else if (player.playerChoice == "rock" || player.playerChoice == "paper" || player.playerChoice == "scissors") && ai.aiItemDecision != "defend" {
		ai.fortressHealth -= (10 + 5*player.successiveVictories)
		// if player did not choose RPS or defend, and ai chose RPS
	} else if player.playerChoice != "defend" && ai.aiChoice >= 5 {
		player.fortressHealth -= (10 + 5*ai.successiveVictories)
	}

	// check player item usage
	if player.playerChoice == "fire_bomb" {
		ai.tickDamage = 5
		player.items[0] -= 1
	} else if player.playerChoice == "cannon_ball" {
		ai.fortressHealth -= 20
		player.items[1] -= 1
	} else if player.playerChoice == "smoke_screen" {
		ai.isSmoked = true
		player.items[2] -= 1
	}

	// check ai item usage
	if ai.aiChoice < 5 {
		if ai.aiItemDecision == "fire_bomb" {
			player.tickDamage = 5
			ai.items[0] -= 1
		} else if ai.aiItemDecision == "cannon_ball" {
			player.fortressHealth -= 20
			ai.items[1] -= 1
		} else if ai.aiItemDecision == "smoke_screen" {
			player.isSmoked = true
			ai.items[2] -= 1
		}
	}
}

func adjustStatus(player *character, ai *character) {
	// reset defending
	player.isDefending = false
	ai.isDefending = false

	// reset is smoked if no smoke screen was used during round
	if player.isSmoked && ai.aiItemDecision != "smoke_screen" {
		player.isSmoked = false
	}
	if ai.isSmoked && player.playerChoice != "smoke_screen" {
		ai.isSmoked = false
	}

	// apply tick damage
	if player.tickDamage > 0 {
		player.fortressHealth -= 5
		player.tickDamage -= 1
	}
	if ai.tickDamage > 0 {
		ai.fortressHealth -= 5
		ai.tickDamage -= 1
	}

	// award random item to player or ai
	item := rand.IntN(3)
	which := rand.IntN(2)
	if which == 0 { // if player chosen
		if item == 0 {
			player.items[0] += 1
		} else if item == 1 {
			player.items[1] += 1
		} else {
			player.items[2] += 1
		}
	} else { // if ai chosen
		if item == 0 {
			ai.items[0] += 1
		} else if item == 1 {
			ai.items[1] += 1
		} else {
			ai.items[2] += 1
		}
	}
}

func printStatus(which int, cc character) {
	if which == 0 {
		fmt.Println("Player Status")
	} else {
		fmt.Println("AI Status")
	}
	fmt.Println("-------------")
	fmt.Println("Fortress Health:", cc.fortressHealth)
	fmt.Println("Fire Bombs:", cc.items[0])
	fmt.Println("Cannon Balls:", cc.items[1])
	fmt.Println("Smoke Screens:", cc.items[2])
	fmt.Println("Is Smoked:", cc.isSmoked)
	fmt.Println("Tick Damage Rounds Remaining:", cc.tickDamage)
	fmt.Println("Successive Victories:", cc.successiveVictories)
	fmt.Println("")
}

func playerMove(player *character, ai *character) {
	for {
		fmt.Scanln(&player.playerChoice)
		if player.playerChoice == "rock" || player.playerChoice == "paper" || player.playerChoice == "scissors" {
			if player.isSmoked {
				fmt.Println("Cannot attack, you are smoked. Choose again.")
				continue
			}
			ai.aiChoice = rand.IntN(10)
			break
		} else if player.playerChoice == "defend" {
			ai.aiChoice = rand.IntN(10)
			player.isDefending = true
			break
		} else if player.playerChoice == "items" {
			if player.items[0] > 0 || player.items[1] > 0 || player.items[2] > 0 {
				fmt.Println("Choose item:")
				fmt.Scanln(&player.playerChoice)
				if player.playerChoice == "fire_bomb" || player.playerChoice == "cannon_ball" || player.playerChoice == "smoke_screen" {
					ai.aiChoice = rand.IntN(10)
					break
				} else {
					fmt.Println("Invalid Choice.")
					fmt.Println("Type in 'fire_bomb', 'cannon_ball', or 'smoke_screen' to choose.")
					fmt.Println("Returning to basic moves. Type 'items' to try again.")
					continue
				}
			} else {
				fmt.Println("No items available to you.")
				continue
			}
		} else if player.playerChoice == "exit" {
			fmt.Println("Ending Game...")
			break
		} else if player.playerChoice == "manual" {
			manual()
			continue
		} else {
			fmt.Println("Invalid Choice.")
			fmt.Println("Type in 'rock', 'paper', 'scissors', 'defend', or 'items' to choose.")
			fmt.Println("Type in 'exit' to end the game.")
			continue
		}
	}
}

func aiMove(ai *character) {
	if ai.isSmoked {
		ai.aiChoice = rand.IntN(5)
	}
	num := rand.IntN(3) // use to decide rps or item choice

	// items
	if ai.aiChoice <= 2 {
		if ai.items[0] > 0 && ai.aiChoice == 0 { // ai choice options
			ai.aiItemDecision = "fire_bomb"
		} else if ai.items[1] > 0 && ai.aiChoice == 1 {
			ai.aiItemDecision = "cannon_ball"
		} else if ai.items[2] > 0 && ai.aiChoice == 2 {
			ai.aiItemDecision = "smoke_screen"
		} else if ai.items[0] > 0 && num == 0 { // num options
			ai.aiItemDecision = "fire_bomb"
		} else if ai.items[1] > 0 && num == 1 {
			ai.aiItemDecision = "cannon_ball"
		} else if ai.items[2] > 0 && num == 2 {
			ai.aiItemDecision = "smoke_screen"
		} else {
			ai.aiChoice = 3
		}
		// basically, if the preferred item isn't available
		// then select the first available item, otherwise
		// choose to defend
	}

	// separate if statement in case ai has no items
	// defend
	if ai.aiChoice > 2 && ai.aiChoice <= 4 {
		ai.aiItemDecision = "defend"
		ai.isDefending = true
	} else if ai.aiChoice >= 5 && !ai.isSmoked { // RPS
		if ai.aiRPSWeight == 3 { // attempt to predict player move
			if ai.aiRPSDecision == "rock" {
				ai.aiRPSDecision = "scissors"
			} else if ai.aiRPSDecision == "paper" {
				ai.aiRPSDecision = "rock"
			} else if ai.aiRPSDecision == "scissors" {
				ai.aiRPSDecision = "paper"
			}
		} else if ai.aiRPSWeight == 2 {
			// do nothing, keep same decision from previous RPS
		} else if num == 0 { // random decision
			ai.aiRPSDecision = "rock"
		} else if num == 1 {
			ai.aiRPSDecision = "paper"
		} else if num == 2 {
			ai.aiRPSDecision = "scissors"
		}
	}
}
