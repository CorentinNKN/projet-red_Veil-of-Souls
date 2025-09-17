package mapgame

import (
	"fmt"
	"main/character"
	"main/utils"
	"math/rand"
)

type Room struct {
	Name        string
	Grid        [][]string
	Connections map[string]*Room
}

func initRooms() *Room {
	s1 := &Room{
		Name: "Salle 1 (entrée)",
		Grid: [][]string{
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", "😈", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", "😈", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", "😈", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
		},
		Connections: make(map[string]*Room),
	}

	s2 := &Room{
		Name: "Salle 2 (plus difficile)",
		Grid: [][]string{
			{"😈", ".", ".", ".", ".", ".", ".", "😈"},
			{".", ".", "😈", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", "😈", ".", ".", ".", "."},
			{"😈", ".", ".", ".", "😈", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{"😈", ".", ".", ".", ".", ".", ".", "😈"},
		},
		Connections: make(map[string]*Room),
	}

	s3 := &Room{
		Name: "Salle 3 (boss léger)",
		Grid: [][]string{
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", "😈", ".", ".", ".", ".", ".", "."},
			{".", ".", "😈", "👹", "😈", ".", ".", "."},
			{".", ".", ".", "😈", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "."},
		},
		Connections: make(map[string]*Room),
	}

	// Relier les salles
	s1.Connections["nord"] = s2
	s2.Connections["sud"] = s1
	s2.Connections["est"] = s3
	s3.Connections["ouest"] = s2

	return s1
}

// ExploreDungeon : parcourt le donjon
func ExploreDungeon(c *character.Character) {
	currentRoom := initRooms()

	for {
		fmt.Printf("\n=== %s ===\n", currentRoom.Name)
		playRoom(c, currentRoom.Grid)

		// vérifier si le joueur est mort définitif
		if c.CurrentHP <= 0 {
			fmt.Println("💀 Vous êtes mort. Fin du jeu.")
			return
		}

		// Choisir sortie
		if len(currentRoom.Connections) == 0 {
			fmt.Println("✔ Vous avez nettoyé la dernière salle, bravo !")
			return
		}

		fmt.Println("\nSorties disponibles :")
		for dir := range currentRoom.Connections {
			fmt.Println("-", dir)
		}
		fmt.Print("Choisissez une direction : ")
		choice := utils.AskChoice()

		if next, ok := currentRoom.Connections[choice]; ok {
			currentRoom = next
		} else {
			fmt.Println("❌ Direction invalide, vous restez dans la salle.")
		}
	}
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

		// Vérifier ce qu’il y a dans la case
		cell := grid[playerX][playerY]
		if cell == "😈" || cell == "👹" {
			fmt.Printf("⚔️ Un ennemi %s apparaît !\n", cell)

			// Dégâts aléatoires
			damage := rand.Intn(20) + 10
			c.CurrentHP -= damage
			fmt.Printf("Vous subissez %d PV de dégâts (%d/%d).\n", damage, c.CurrentHP, c.MaxHP)

			if character.IsDead(c) {
				fmt.Println("⚡ Vous avez été ressuscité à 50% de vos PV.")
			}

			// Ennemi battu → case vidée
			grid[playerX][playerY] = "."
		}

		// ✅ Si plus aucun ennemi → sortie automatique
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
