package mapgame

import (
	"fmt"
	"main/character"
	"math/rand"
	"time"
)

// Dimensions de la map
const (
	width  = 10
	height = 6
)

type Map struct {
	grid    [][]string
	playerX int
	playerY int
}

// Crée une map avec joueur et ennemis
func InitMap() *Map {
	m := &Map{
		grid:    make([][]string, height),
		playerX: width / 2,
		playerY: height / 2,
	}

	for i := range m.grid {
		m.grid[i] = make([]string, width)
		for j := range m.grid[i] {
			m.grid[i][j] = "⬜" // case vide
		}
	}

	// Place joueur
	m.grid[m.playerY][m.playerX] = "🥷"

	// Place ennemis
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		ex := rand.Intn(width)
		ey := rand.Intn(height)
		if ex != m.playerX || ey != m.playerY {
			m.grid[ey][ex] = "😈"
		}
	}

	return m
}

// Affiche la map
func (m *Map) Display() {
	fmt.Println("===== MAP =====")
	for _, row := range m.grid {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
	fmt.Println("================")
}

// Déplace le joueur
func (m *Map) Move(dir string) {
	// efface l'ancienne position
	m.grid[m.playerY][m.playerX] = "⬜"

	switch dir {
	case "z":
		if m.playerY > 0 {
			m.playerY--
		}
	case "s":
		if m.playerY < height-1 {
			m.playerY++
		}
	case "q":
		if m.playerX > 0 {
			m.playerX--
		}
	case "d":
		if m.playerX < width-1 {
			m.playerX++
		}
	}

	// Vérifie ce qu'il y a sur la case
	if m.grid[m.playerY][m.playerX] == "😈" {
		fmt.Println("⚔️  Un combat commence contre un ennemi ! 😈")
	}

	// place le joueur
	m.grid[m.playerY][m.playerX] = "🥷"
}

func StartExploration(c *character.Character) {
	fmt.Println("Exploration de la carte... (fonctionnalité à implémenter)")
}
