package mapgame

import (
	"fmt"
	"main/character"
	"main/utils"
	"math/rand"
)

var rooms = [][][]string{
	{ // Salle 1
		{".", ".", ".", "."},
		{".", "😈", ".", "."},
		{".", ".", ".", "😈"},
		{".", ".", ".", "."},
	},
	{ // Salle 2 (plus difficile)
		{"😈", ".", ".", "😈"},
		{".", "😈", ".", "."},
		{".", ".", ".", "."},
		{"😈", ".", "😈", "."},
	},
	{ // Salle 3 (boss léger)
		{".", ".", ".", "."},
		{".", "😈", "😈", "."},
		{".", "😈", "👹", "😈"},
		{".", ".", ".", "."},
	},
}

// ExploreRooms : parcourt les salles
func ExploreRooms(c *character.Character) {
	for i, room := range rooms {
		fmt.Printf("\n=== Salle %d ===\n", i+1)
		playRoom(c, room)
		// si après la salle le joueur a 0 ou moins => IsDead a été appelé dans playRoom, continue
	}
	fmt.Println("✔ Vous avez terminé toutes les salles disponibles.")
}

func playRoom(c *character.Character, grid [][]string) {
	playerX, playerY := 0, 0

	for {
		displayMap(playerX, playerY, grid)
		fmt.Println("Déplacez-vous (z: haut, s: bas, q: gauche, d: droite, r: quitter la salle)")
		choice := utils.AskChoice()

		switch choice {
		case "z":
			if playerX > 0 {
				playerX--
			}
		case "s":
			if playerX < len(grid)-1 {
				playerX++
			}
		case "q":
			if playerY > 0 {
				playerY--
			}
		case "d":
			if playerY < len(grid[0])-1 {
				playerY++
			}
		case "r":
			return
		default:
			fmt.Println("Mauvais choix.")
		}

		cell := grid[playerX][playerY]
		if cell == "😈" || cell == "👹" {
			// combat simplifié : on subit des dégâts aléatoires, ennemi supprimé après
			fmt.Printf("⚔️ Un ennemi %s apparaît !\n", cell)
			damage := rand.Intn(20) + 10
			c.CurrentHP -= damage
			fmt.Printf("Vous subissez %d PV de dégâts (%d/%d).\n", damage, c.CurrentHP, c.MaxHP)
			// vérifier mort (IsDead gère la résurrection)
			if character.IsDead(c) {
				fmt.Println("⚡ Vous avez été ressuscité à 50% de vos PV.")
			}
			// retirer l'ennemi
			grid[playerX][playerY] = "."
		}

		// vérifier si la salle est nettoyée
		if isRoomCleared(grid) {
			fmt.Println("✔ Salle nettoyée !")
			return
		}
	}
}

func displayMap(playerX, playerY int, grid [][]string) {
	fmt.Println("\n--- Carte ---")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == playerX && j == playerY {
				fmt.Print("🥷 ")
			} else {
				fmt.Print(grid[i][j] + " ")
			}
		}
		fmt.Println()
	}
}

func isRoomCleared(grid [][]string) bool {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "😈" || grid[i][j] == "👹" {
				return false
			}
		}
	}
	return true
}
