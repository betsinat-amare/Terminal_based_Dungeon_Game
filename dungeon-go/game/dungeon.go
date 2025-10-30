package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Room represents a single room in the dungeon
type Room struct {
	X        int
	Y        int
	Visited  bool
	HasEnemy bool
	HasItem  bool
	IsExit   bool
	// you could expand with Enemy pointer here, but we spawn enemies on-the-fly
}

// Dungeon represents the full map
type Dungeon struct {
	Width, Height    int
	Grid             [][]Room
	PlayerX, PlayerY int
}

// NewDungeon generates a new random dungeon
func NewDungeon(width, height int) *Dungeon {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]Room, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]Room, width)
		for x := 0; x < width; x++ {
			grid[y][x] = Room{
				X:        x,
				Y:        y,
				Visited:  false,
				HasEnemy: rand.Float32() < 0.2,  // 20% chance of enemy
				HasItem:  rand.Float32() < 0.15, // 15% chance of item
				IsExit:   false,
			}
		}
	}

	// Randomly pick an exit room
	exitX := rand.Intn(width)
	exitY := rand.Intn(height)
	grid[exitY][exitX].IsExit = true

	d := &Dungeon{
		Width:  width,
		Height: height,
		Grid:   grid,
	}

	// Place player at a random start position
	d.PlayerX = rand.Intn(width)
	d.PlayerY = rand.Intn(height)
	grid[d.PlayerY][d.PlayerX].Visited = true

	return d
}

// DisplayCurrentRoom shows info about where the player is
func (d *Dungeon) DisplayCurrentRoom() {
	room := d.Grid[d.PlayerY][d.PlayerX]
	fmt.Printf("\n📍 You are in room (%d, %d)\n", room.X, room.Y)

	if room.IsExit {
		fmt.Println("✨ You found the exit! You win!")
		return
	}

	if room.HasEnemy {
		fmt.Println("⚔️  You sense danger... an enemy lurks here!")
	} else {
		fmt.Println("The room is quiet.")
	}

	if room.HasItem {
		fmt.Println("💎 There’s something shiny on the floor.")
	}
}

// MovePlayer moves the player in a direction (if valid)
func (d *Dungeon) MovePlayer(direction string) {
	newX, newY := d.PlayerX, d.PlayerY

	switch direction {
	case "n":
		newY--
	case "s":
		newY++
	case "e":
		newX++
	case "w":
		newX--
	default:
		fmt.Println("Invalid direction! Use n/s/e/w.")
		return
	}

	if newX < 0 || newY < 0 || newX >= d.Width || newY >= d.Height {
		fmt.Println("🚧 You hit a wall. Try another direction.")
		return
	}

	d.PlayerX = newX
	d.PlayerY = newY
	d.Grid[newY][newX].Visited = true

	d.DisplayCurrentRoom()
}

// CheckForEnemy returns an Enemy pointer if the current room has an enemy.
// It also returns a boolean indicating whether there was an enemy to begin with.
func (d *Dungeon) CheckForEnemy() (*Enemy, bool) {
	room := &d.Grid[d.PlayerY][d.PlayerX]
	if !room.HasEnemy {
		return nil, false
	}
	// Spawn an enemy based on a difficulty factor (e.g., distance from start)
	// compute depth as Manhattan distance from origin (0,0) as a simple heuristic
	depthFactor := abs(d.PlayerX) + abs(d.PlayerY)
	enemy := NewEnemy(depthFactor)
	return enemy, true
}

// ClearEnemy clears the HasEnemy flag for current room (call after enemy defeated)
func (d *Dungeon) ClearEnemy() {
	d.Grid[d.PlayerY][d.PlayerX].HasEnemy = false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
