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
}
