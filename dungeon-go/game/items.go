package game

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Item interface
type Item interface {
	Name() string
	Use(*Player)
}

// Potion struct
type Potion struct {
	Heal int
}

func (p Potion) Name() string { return fmt.Sprintf("Health Potion (+%d HP)", p.Heal) }
func (p Potion) Use(player *Player) {
	player.HP += p.Heal
	fmt.Printf("💖 You drink the potion and restore %d HP. (HP: %d)\n", p.Heal, player.HP)
}

// Weapon struct
type Weapon struct {
	NameStr string
	Attack  int
}

func (w Weapon) Name() string { return fmt.Sprintf("%s (+%d ATK)", w.NameStr, w.Attack) }
func (w Weapon) Use(player *Player) {
	player.Attack += w.Attack
	fmt.Printf("⚔️ You equip %s! (Attack: %d)\n", w.Name(), player.Attack)
}

// --- JSON Wrapper ---

type SaveItem struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// Converts Item to SaveItem
func WrapItem(item Item) (SaveItem, error) {
	var data []byte
	var err error
	switch v := item.(type) {
	case Potion:
		data, err = json.Marshal(v)
		if err != nil {
			return SaveItem{}, err
		}
		return SaveItem{Type: "Potion", Data: data}, nil
	case Weapon:
		data, err = json.Marshal(v)
		if err != nil {
			return SaveItem{}, err
		}
		return SaveItem{Type: "Weapon", Data: data}, nil
	default:
		return SaveItem{}, fmt.Errorf("unknown item type")
	}
}

// Converts SaveItem back to Item
func UnwrapItem(s SaveItem) (Item, error) {
	switch s.Type {
	case "Potion":
		var p Potion
		err := json.Unmarshal(s.Data, &p)
		return p, err
	case "Weapon":
		var w Weapon
		err := json.Unmarshal(s.Data, &w)
		return w, err
	default:
		return nil, fmt.Errorf("unknown save item type: %s", s.Type)
	}
}
