package game

import (
	"fmt"
	"math/rand"
)

// Enemy represents a monster in the dungeon
type Enemy struct {
	Name    string
	HP      int
	Attack  int
	Defense int
	Level   int
}

// NewEnemy creates an enemy based on a simple type selection
func NewEnemy(depthFactor int) *Enemy {
	// depthFactor can be used later to scale difficulty by dungeon depth/level
	roll := rand.Intn(100)
	if roll < 60 {
		// Common: Goblin
		return &Enemy{
			Name:    "Goblin",
			HP:      20 + depthFactor*2,
			Attack:  5 + depthFactor,
			Defense: 1 + depthFactor/2,
			Level:   1 + depthFactor/2,
		}
	} else if roll < 90 {
		// Uncommon: Orc
		return &Enemy{
			Name:    "Orc",
			HP:      35 + depthFactor*3,
			Attack:  8 + depthFactor*2,
			Defense: 3 + depthFactor/2,
			Level:   2 + depthFactor/2,
		}
	}
	// Rare: Worg (or mini-boss)
	return &Enemy{
		Name:    "Worg",
		HP:      55 + depthFactor*4,
		Attack:  12 + depthFactor*3,
		Defense: 5 + depthFactor,
		Level:   3 + depthFactor,
	}
}

// IsAlive checks if enemy still has HP
func (e *Enemy) IsAlive() bool {
	return e.HP > 0
}

// AttackPlayer performs an attack on the player and returns damage value
func (e *Enemy) AttackPlayer(p *Player) int {
	raw := rand.Intn(e.Attack/2+1) + e.Attack/2 // between Attack/2 .. Attack
	// player defense handled in Player.TakeDamage
	return raw
}

// PossibleLoot returns a random loot string (or empty == nothing)
func (e *Enemy) PossibleLoot() Item {
	chance := rand.Intn(100)
	if chance < 40 {
		return nil // no loot
	}
	return GenerateRandomItem()
}

func (e *Enemy) String() string {
	return fmt.Sprintf("%s (HP:%d ATK:%d DEF:%d)", e.Name, e.HP, e.Attack, e.Defense)
}
