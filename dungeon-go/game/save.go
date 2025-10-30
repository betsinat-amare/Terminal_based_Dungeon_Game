package game

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveData defines the overall game state we serialize
type SaveData struct {
	Player  Player
	Dungeon Dungeon
}

// SaveGame writes the current player and dungeon state to a file
func SaveGame(filename string, player *Player, dungeon *Dungeon) error {
	data := SaveData{
		Player:  *player,
		Dungeon: *dungeon,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create save file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("could not encode save data: %v", err)
	}

	fmt.Printf("💾 Game saved to %s!\n", filename)
	return nil
}

// LoadGame loads the saved game state from a file
func LoadGame(filename string) (*Player, *Dungeon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open save file: %v", err)
	}
	defer file.Close()

	var data SaveData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, nil, fmt.Errorf("could not decode save data: %v", err)
	}

	fmt.Printf("📂 Game loaded from %s!\n", filename)
	return &data.Player, &data.Dungeon, nil
}
