package mapgame

import (
	"encoding/json"
	"fmt"
	"main/character"
	"main/inventory"
	"main/utils"
	"math/rand"
	"os"
)

// Room représente une salle du donjon
type Room struct {
	Name        string
	Grid        [][]string
	Connections map[string]*Room
	IsFinal     bool
}

// --- Création des salles ---
func initRooms() *Room {
	var previous *Room
	var first *Room

	// Générer 10 salles + boss (11 salles au total)
	for i := 1; i <= 11; i++ {
		room := &Room{
			Name:        fmt.Sprintf("Salle %d", i),
			Grid:        generateRoom(i),
			Connections: make(map[string]*Room),
			IsFinal:     (i == 11),
		}

		if previous != nil {
			previous.Connections["est"] = room
			room.Connections["ouest"] = previous
		} else {
			first = room
		}
		previous = room
	}
	return first
}

// Génère une grille de taille fixe (8x8) et place des ennemis selon le niveau
func generateRoom(level int) [][]string {
	size := 8
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Nombre d’ennemis selon la difficulté (plus de niveau => plus d'ennemis)
	numEnemies := level + rand.Intn(3)
	for k := 0; k < numEnemies; k++ {
		// trouver une case vide pour poser l'ennemi (évite écraser)
		for {
			x, y := rand.Intn(size), rand.Intn(size)
			if grid[x][y] == "." {
				if level == 11 {
					grid[x][y] = "👹" // boss final(s)
				} else if rand.Intn(5) == 0 {
					grid[x][y] = "👹" // mini-boss rare
				} else {
					grid[x][y] = "😈"
				}
				break
			}
		}
	}
	return grid
}

// --- Exploration du donjon ---
func ExploreDungeon(c *character.Character) {
	currentRoom := LoadGame()
	if currentRoom == nil {
		currentRoom = initRooms()
	}

	for {
		fmt.Printf("\n=== %s ===\n", currentRoom.Name)
		playRoom(c, currentRoom)

		// joueur mort en dehors de la résurrection gérée par character.IsDead
		if c.CurrentHP <= 0 {
			fmt.Println("💀 Vous êtes mort. Fin du jeu.")
			_ = os.Remove("save.json") // supprimer sauvegarde
			return
		}

		// Boss final terminé → victoire
		if currentRoom.IsFinal && isRoomCleared(currentRoom.Grid) {
			fmt.Println("\n🎉🎉 YOU WIN! 🎉🎉")
			_ = os.Remove("save.json") // reset save
			return
		}

		// si pas de connexions, fin
		if len(currentRoom.Connections) == 0 {
			fmt.Println("✔ Vous avez nettoyé la dernière salle, bravo !")
			_ = os.Remove("save.json")
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
			SaveGame(currentRoom) // sauvegarder la salle courante
		} else {
			fmt.Println("❌ Direction invalide, vous restez dans la salle.")
		}
	}
}

// --- Jouer une salle ---
// NOTE : playRoom prend maintenant *Room (pas [][]string) -> évite les erreurs de type
func playRoom(c *character.Character, room *Room) {
	grid := room.Grid
	playerX, playerY := 0, 0

	for {
		displayMap(playerX, playerY, grid)
		fmt.Println("Déplacez-vous (z: haut, s: bas, q: gauche, d: droite, i: inventaire, r: quitter la salle)")
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
		case "i":
			// accès à l'inventaire pendant l'exploration
			inventory.AccessInventory(c)
		case "r":
			return
		default:
			fmt.Println("Mauvais choix.")
		}

		// Combat si ennemi
		cell := grid[playerX][playerY]
		if cell == "😈" || cell == "👹" {
			fmt.Printf("⚔️ Un ennemi %s apparaît !\n", cell)
			damage := rand.Intn(20) + 10
			c.CurrentHP -= damage
			fmt.Printf("Vous subissez %d PV de dégâts (%d/%d).\n", damage, c.CurrentHP, c.MaxHP)

			if character.IsDead(c) {
				fmt.Println("⚡ Vous avez été ressuscité à 50% de vos PV.")
			}

			// retirer l'ennemi
			grid[playerX][playerY] = "."
		}

		// Salle terminée -> sortie immédiate (plus besoin de tourner dans la salle vide)
		if isRoomCleared(grid) {
			fmt.Println("✔ Salle nettoyée !")
			return
		}
	}
}

// --- Affichage ---
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

// Vérifie si salle nettoyée
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

// --- Sauvegarde simple (sauve juste le nom de la salle) ---
type SaveData struct {
	RoomName string `json:"room_name"`
}

func SaveGame(room *Room) {
	data := SaveData{RoomName: room.Name}
	file, err := os.Create("save.json")
	if err != nil {
		// on ne plante pas le jeu pour une erreur de sauvegarde
		fmt.Println("⚠️ Impossible d'écrire la sauvegarde :", err)
		return
	}
	defer file.Close()
	_ = json.NewEncoder(file).Encode(data)
}

// Recréé le donjon et cherche la salle par nom (BFS pour être sûr)
func LoadGame() *Room {
	file, err := os.Open("save.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	var data SaveData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil
	}

	// Reconstruire le donjon et retrouver la salle sauvegardée
	first := initRooms()
	// BFS pour trouver la salle qui a le nom sauvegardé
	seen := map[*Room]bool{}
	queue := []*Room{first}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}
		if seen[cur] {
			continue
		}
		seen[cur] = true
		if cur.Name == data.RoomName {
			return cur
		}
		for _, next := range cur.Connections {
			if !seen[next] {
				queue = append(queue, next)
			}
		}
	}
	// si non trouvé, retourne la première salle
	return first
}
