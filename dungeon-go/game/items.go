package game

import (
	"fmt"
	"math/rand"
)

// Item is a general interface for usable or collectible objects
type Item interface {
	Name() string
	Use(p *Player)
}

// Potion restores HP when used
type Potion struct {
	HealAmount int
}

func (potion Potion) Name() string {
	return fmt.Sprintf("Health Potion (+%d HP)", potion.HealAmount)
}

func (potion Potion) Use(p *Player) {
	p.HP += potion.HealAmount
	if p.HP > 100 {
		p.HP = 100
	}
	fmt.Printf("💖 You drink the potion and restore %d HP. (HP: %d)\n", potion.HealAmount, p.HP)
}

// Weapon increases attack stat
type Weapon struct {
	WeaponName string
	Bonus      int
}

func (w Weapon) Name() string {
	return fmt.Sprintf("%s (+%d Attack)", w.WeaponName, w.Bonus)
}

func (w Weapon) Use(p *Player) {
	p.Attack += w.Bonus
	fmt.Printf("⚔️  You equip the %s! Attack increased by %d (Now: %d)\n", w.WeaponName, w.Bonus, p.Attack)
}

// GenerateRandomItem returns a random item instance
func GenerateRandomItem() Item {
	roll := rand.Intn(100)
	if roll < 60 {
		return Potion{HealAmount: 20 + rand.Intn(10)}
	}
	names := []string{"Iron Sword", "Silver Dagger", "Magic Spear"}
	return Weapon{WeaponName: names[rand.Intn(len(names))], Bonus: 3 + rand.Intn(4)}
}
