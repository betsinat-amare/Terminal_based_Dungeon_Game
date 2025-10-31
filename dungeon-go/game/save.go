package game

import (
	"encoding/json"
	"fmt"
	"os"
)

type SaveData struct {
	Player  Player     `json:"player"`
	Dungeon Dungeon    `json:"dungeon"`
	Items   []SaveItem `json:"items"` // serialized inventory
}

// SaveGame writes player, dungeon, and inventory
func SaveGame(filename string, player *Player, dungeon *Dungeon) error {
	items := []SaveItem{}
	for _, item := range player.Inventory {
		wrapped, err := WrapItem(item)
		if err != nil {
			return err
		}
		items = append(items, wrapped)
	}

	data := SaveData{
		Player:  *player,
		Dungeon: *dungeon,
		Items:   items,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create save file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("could not encode save data: %v", err)
	}

	fmt.Printf("💾 Game saved to %s!\n", filename)
	return nil
}

// LoadGame restores player, dungeon, and inventory
func LoadGame(filename string) (*Player, *Dungeon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open save file: %v", err)
	}
	defer file.Close()

	var data SaveData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, nil, fmt.Errorf("could not decode save data: %v", err)
	}

	// Unwrap items
	inventory := []Item{}
	for _, s := range data.Items {
		item, err := UnwrapItem(s)
		if err != nil {
			return nil, nil, fmt.Errorf("could not restore item: %v", err)
		}
		inventory = append(inventory, item)
	}
	data.Player.Inventory = inventory

	fmt.Printf("📂 Game loaded from %s!\n", filename)
	return &data.Player, &data.Dungeon, nil
}
