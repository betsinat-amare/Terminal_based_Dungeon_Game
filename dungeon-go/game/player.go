package game

import "fmt"

// Player represents the adventurer exploring the dungeon
type Player struct {
	Name      string
	HP        int
	Attack    int
	Defense   int
	Inventory []string
}

// NewPlayer creates and returns a new Player
func NewPlayer(name string) Player {
	return Player{
		Name:      name,
		HP:        100,
		Attack:    10,
		Defense:   5,
		Inventory: []string{},
	}
}

// ShowStats displays the player's current stats
func (p *Player) ShowStats() {
	fmt.Printf("\n👤 %s's Stats:\n", p.Name)
	fmt.Printf("HP: %d\n", p.HP)
	fmt.Printf("Attack: %d\n", p.Attack)
	fmt.Printf("Defense: %d\n", p.Defense)
	fmt.Printf("Inventory: %v\n\n", p.Inventory)
}
