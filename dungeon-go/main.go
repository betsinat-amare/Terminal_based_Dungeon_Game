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
		// 🧟 Check for enemy in the current room BEFORE player actions
		if enemy, ok := dungeon.CheckForEnemy(); ok {
			fmt.Printf("\nYou encounter a %s!\n", enemy.String())

			for {
				fmt.Print("Do you want to (f)ight or (r)un? ")
				choiceRaw, _ := reader.ReadString('\n')
				choice := strings.TrimSpace(choiceRaw)

				if choice == "r" {
					fmt.Println("You attempt to flee...")
					// Try to move in a random direction to simulate escaping
					for _, dir := range []string{"s", "n", "e", "w"} {
						prevX, prevY := dungeon.PlayerX, dungeon.PlayerY
						dungeon.MovePlayer(dir)
						if dungeon.PlayerX != prevX || dungeon.PlayerY != prevY {
							fmt.Println("You successfully fled to another room!")
							break
						}
					}
					break // end the encounter
				}

				if choice == "f" {
					// ⚔️ Combat loop
					for enemy.IsAlive() && player.IsAlive() {
						dmg := player.AttackEnemy(enemy)
						fmt.Printf("You strike the %s for %d damage! (Enemy HP: %d)\n", enemy.Name, dmg, enemy.HP)

						if !enemy.IsAlive() {
							fmt.Printf("🎉 You defeated the %s!\n", enemy.Name)
							loot := enemy.PossibleLoot()
							if loot != nil {
								fmt.Printf("💎 You found: %s\n", loot.Name())
								player.AddItem(loot)
							} else {
								fmt.Println("No loot this time.")
							}
							dungeon.ClearEnemy()
							break
						}

						// Enemy’s turn
						edmg := enemy.AttackPlayer(player)
						remaining := player.TakeDamage(edmg)
						fmt.Printf("The %s hits you for %d! (Your HP: %d)\n", enemy.Name, edmg, remaining)

						if !player.IsAlive() {
							fmt.Println("💀 You have been defeated... Game over.")
							return
						}
					}
					break // end combat
				}

				fmt.Println("Invalid choice. Type 'f' to fight or 'r' to run.")
			}
		}

		// 🎮 Player action loop
		fmt.Print("\nAction → move (n/s/e/w), use, stats, or q to quit: ")
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		switch input {
		case "q":
			fmt.Println("You have chosen to leave the dungeon.")
			return

		case "stats":
			player.ShowStats()
			continue

		case "use":
			if len(player.Inventory) == 0 {
				fmt.Println("🎒 Your inventory is empty.")
				continue
			}

			fmt.Println("Select an item to use:")
			for i, item := range player.Inventory {
				fmt.Printf("%d) %s\n", i+1, item.Name())
			}

			var choice int
			fmt.Print("Enter number: ")
			fmt.Scanln(&choice)

			if choice < 1 || choice > len(player.Inventory) {
				fmt.Println("Invalid choice.")
				continue
			}

			item := player.Inventory[choice-1]
			item.Use(player)

			// Remove potions after use
			switch item.(type) {
			case game.Potion:
				player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
			}
			continue

		default:
			dungeon.MovePlayer(input)
			room := dungeon.Grid[dungeon.PlayerY][dungeon.PlayerX]

			if room.HasItem {
				item := game.GenerateRandomItem()
				fmt.Printf("💎 You found a %s!\n", item.Name())
				player.AddItem(item)
				room.HasItem = false
			}
		}
	}
}
