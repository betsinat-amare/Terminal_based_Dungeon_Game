package game

import (
	"fmt"
	"math/rand"
)

// Player represents the adventurer exploring the dungeon
type Player struct {
	Name      string
	HP        int
	Attack    int
	Defense   int
	Level     int
	Inventory []Item
}

// NewPlayer creates and returns a new Player
func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		HP:        100,
		Attack:    10,
		Defense:   5,
		Inventory: []Item{},
		Level:     1,
	}
}

// ShowStats displays the player's current stats
func (p *Player) ShowStats() {
	fmt.Printf("\n👤 %s's Stats:\n", p.Name)
	fmt.Printf("HP: %d\n", p.HP)
	fmt.Printf("Attack: %d\n", p.Attack)
	fmt.Printf("Defense: %d\n", p.Defense)
	fmt.Printf("Level: %d\n", p.Level)

	fmt.Println("🎒 Inventory:")
	if len(p.Inventory) == 0 {
		fmt.Println("  (empty)")
	} else {
		for i, item := range p.Inventory {
			fmt.Printf("  %d) %s\n", i+1, item.Name())
		}
	}
	fmt.Println()
}

// AttackEnemy performs an attack on the enemy and returns the damage dealt
func (p *Player) AttackEnemy(e *Enemy) int {
	// randomize damage a bit: base up to Attack
	raw := rand.Intn(p.Attack/2+1) + p.Attack/2 // between Attack/2 .. Attack
	// enemy's defense reduces damage
	damage := raw - e.Defense
	if damage < 1 {
		damage = 1
	}
	e.HP -= damage
	return damage
}

// TakeDamage reduces player's HP by damage and returns remaining HP
func (p *Player) TakeDamage(damage int) int {
	// simple mitigation by defense
	actual := damage - (p.Defense / 2)
	if actual < 1 {
		actual = 1
	}
	p.HP -= actual
	return p.HP
}

// AddItem adds an item to inventory
func (p *Player) AddItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

// IsAlive returns whether player still has HP
func (p *Player) IsAlive() bool {
	return p.HP > 0
}
