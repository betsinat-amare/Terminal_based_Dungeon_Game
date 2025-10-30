package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dungeon-go/game"
)

func main() {
	fmt.Println(game.Cyan + "Welcome to Terminal Dungeon 🏰" + game.Reset)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("1. Start New Game")
	fmt.Println("2. Exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	inputRaw, _ := reader.ReadString('\n')
	input := strings.TrimSpace(inputRaw)

	switch input {
	case "1":
		startGame()
	case "2":
		fmt.Println(game.Yellow + "Goodbye, adventurer!" + game.Reset)
	default:
		fmt.Println(game.Red + "Invalid choice!" + game.Reset)
	}
}

func startGame() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your hero's name: ")
	nameRaw, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameRaw)

	player := game.NewPlayer(name)
	fmt.Printf("\n"+game.Green+"Welcome, %s the Brave!\n"+game.Reset, player.Name)
	player.ShowStats()

	dungeon := game.NewDungeon(5, 5)
	fmt.Println(game.Cyan + "🗺️  A mysterious dungeon appears..." + game.Reset)
	dungeon.DisplayCurrentRoom()

	for {
		// Check for enemy
		if enemy, ok := dungeon.CheckForEnemy(); ok {
			fmt.Printf("\n%sYou encounter a %s!%s\n", game.Red, enemy.String(), game.Reset)

			for {
				fmt.Print("Do you want to (f)ight or (r)un? ")
				choiceRaw, _ := reader.ReadString('\n')
				choice := strings.TrimSpace(choiceRaw)

				if choice == "r" {
					fmt.Println("You attempt to flee...")
					for _, dir := range []string{"s", "n", "e", "w"} {
						prevX, prevY := dungeon.PlayerX, dungeon.PlayerY
						dungeon.MovePlayer(dir)
						if dungeon.PlayerX != prevX || dungeon.PlayerY != prevY {
							fmt.Println("You successfully fled to another room!")
							break
						}
					}
					break
				}

				if choice == "f" {
					for enemy.IsAlive() && player.IsAlive() {
						// Player attacks
						dmg := player.AttackEnemy(enemy)
						fmt.Printf("%sYou strike the %s for %d damage! (Enemy HP: %d)%s\n", game.Green, enemy.Name, dmg, enemy.HP, game.Reset)

						if !enemy.IsAlive() {
							fmt.Printf("%s🎉 You defeated the %s!%s\n", game.Yellow, enemy.Name, game.Reset)
							loot := enemy.PossibleLoot()
							if loot != nil {
								fmt.Printf("%s💎 You found: %s%s\n", game.Cyan, loot.Name(), game.Reset)
								player.AddItem(loot)
							} else {
								fmt.Println("No loot this time.")
							}
							dungeon.ClearEnemy()
							break
						}

						// Enemy attacks
						edmg := enemy.AttackPlayer(player)
						remaining := player.TakeDamage(edmg)
						fmt.Printf("%sThe %s hits you for %d! (Your HP: %d)%s\n", game.Red, enemy.Name, edmg, remaining, game.Reset)

						if !player.IsAlive() {
							fmt.Println(game.Red + "💀 You have been defeated... Game over." + game.Reset)
							return
						}
					}
					break
				}

				fmt.Println("Invalid choice. Type 'f' to fight or 'r' to run.")
			}
		}

		// Player actions
		fmt.Print("\nAction → move (n/s/e/w), use, stats, save, load, or q to quit: ")
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		switch input {
		case "q":
			fmt.Println(game.Yellow + "You have chosen to leave the dungeon." + game.Reset)
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

			switch item.(type) {
			case game.Potion:
				player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
			}
			continue

		case "save":
			err := game.SaveGame("save.json", player, dungeon)
			if err != nil {
				fmt.Println(game.Red+"❌ Error saving game:", err, game.Reset)
			}
			continue

		case "load":
			loadedPlayer, loadedDungeon, err := game.LoadGame("save.json")
			if err != nil {
				fmt.Println(game.Red+"❌ Error loading game:", err, game.Reset)
				continue
			}
			player = loadedPlayer
			dungeon = loadedDungeon
			fmt.Println(game.Green + "✅ Game state restored." + game.Reset)
			dungeon.DisplayCurrentRoom()
			continue

		default:
			dungeon.MovePlayer(input)
			room := dungeon.Grid[dungeon.PlayerY][dungeon.PlayerX]

			if room.HasItem {
				item := game.GenerateRandomItem()
				fmt.Printf("%s💎 You found a %s%s\n", game.Cyan, item.Name(), game.Reset)
				player.AddItem(item)
				room.HasItem = false
			}
		}
	}
}
