package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dungeon-go/game"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your hero's name: ")
	nameRaw, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameRaw)

	player := game.NewPlayer(name)
	fmt.Printf("\nWelcome, %s the Brave!\n", player.Name)
	player.ShowStats()

	dungeon := game.NewDungeon(5, 5)
	fmt.Println("🗺️  A mysterious dungeon appears...")
	dungeon.DisplayCurrentRoom()

	for {
		// Check for enemy in the current room BEFORE prompting movement
		if enemy, ok := dungeon.CheckForEnemy(); ok {
			fmt.Printf("\nYou encounter a %s!\n", enemy.String())
			// Fight or flee prompt
			for {
				fmt.Print("Do you want to (f)ight or (r)un? ")
				choiceRaw, _ := reader.ReadString('\n')
				choice := strings.TrimSpace(choiceRaw)

				if choice == "r" {
					// Running: move to a random adjacent room if possible
					fmt.Println("You attempt to run...")
					// naive flee: just try to move south then north then east then west
					// (could be improved later)
					tried := false
					for _, dir := range []string{"s", "n", "e", "w"} {
						// attempt move, but if wall, MovePlayer will print wall
						prevX, prevY := dungeon.PlayerX, dungeon.PlayerY
						dungeon.MovePlayer(dir)
						// if position changed, we successfully fled
						if dungeon.PlayerX != prevX || dungeon.PlayerY != prevY {
							fmt.Println("You fled to another room.")
							tried = true
							break
						}
					}
					if tried {
						break // exit the fight prompt loop (we fled)
					}
					// couldn't flee; continue fight prompt
					fmt.Println("You couldn't find an escape route!")
					continue
				}

				if choice == "f" {
					// Combat loop
					for enemy.IsAlive() && player.IsAlive() {
						// player's turn
						dmg := player.AttackEnemy(enemy)
						fmt.Printf("You strike the %s for %d damage. (Enemy HP: %d)\n", enemy.Name, dmg, enemy.HP)
						if !enemy.IsAlive() {
							fmt.Printf("You defeated the %s!\n", enemy.Name)
							// loot
							loot := enemy.PossibleLoot()
							if loot != "" {
								fmt.Printf("You found: %s\n", loot)
								player.AddItem(loot)
							} else {
								fmt.Println("No loot this time.")
							}
							// clear enemy from room
							dungeon.ClearEnemy()
							break
						}

						// enemy's turn
						edmg := enemy.AttackPlayer(player)
						remaining := player.TakeDamage(edmg)
						fmt.Printf("The %s hits you for %d. (Your HP: %d)\n", enemy.Name, edmg, remaining)
						if !player.IsAlive() {
							fmt.Println("You have been defeated... Game over.")
							return
						}
					}
					// combat finished (victory or death). break fight prompt
					break
				}

				fmt.Println("Invalid choice. Enter 'f' to fight or 'r' to run.")
			} // end fight-or-run prompt
		} // end if encounter

		// normal movement prompt
		fmt.Print("\nMove (n/s/e/w) or q to quit, or 'stats' to view stats: ")
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		if input == "q" {
			fmt.Println("You have chosen to leave the dungeon.")
			break
		}
		if input == "stats" {
			player.ShowStats()
			continue
		}
		// if user typed nothing (enter), skip
		if input == "" {
			continue
		}
		dungeon.MovePlayer(input)
	}
}
