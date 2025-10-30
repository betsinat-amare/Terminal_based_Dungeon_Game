package main

import (
	"dungeon-go/game"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Terminal Dungeon 🏰")
	fmt.Println("------------------------------")
	fmt.Println("1. Start New Game")
	fmt.Println("2. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		startGame()
	} else {
		fmt.Println("Goodbye, adventurer!")
	}
}

func startGame() {
	var name string
	fmt.Print("Enter your hero's name: ")
	fmt.Scanln(&name)

	player := game.NewPlayer(name)
	fmt.Printf("\nWelcome, %s the Brave!\n", player.Name)
	player.ShowStats()

	dungeon := game.NewDungeon(5, 5)
	fmt.Println("🗺️  A mysterious dungeon appears...")
	dungeon.DisplayCurrentRoom()

	for {
		fmt.Print("\nMove (n/s/e/w) or q to quit: ")
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("You have chosen to leave the dungeon.")
			break
		}

		dungeon.MovePlayer(input)
	}
}
