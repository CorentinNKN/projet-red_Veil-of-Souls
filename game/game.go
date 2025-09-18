package game

import (
	"encoding/json"
	"fmt"
	"main/blacksmith"
	"main/character"
	"main/inventory"
	"main/merchant"
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

func StartGame() {
	fmt.Println("=== Démarrage du jeu ===")
	player := character.CharacterCreation()

	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Accéder au forgeron")
		fmt.Println("5. Explorer le donjon")
		fmt.Println("6. Quitter")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			character.DisplayInfo(&player)
		case "2":
			inventory.AccessInventory(&player)
		case "3":
			merchant.AccessMerchant(&player)
		case "4":
			blacksmith.AccessBlacksmith(&player)
		case "5":
			ExploreDungeon(&player)
		case "6":
			fmt.Println("Merci d’avoir joué à Veil of Souls !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Création des salles
func initRooms() *Room {
	var previous *Room
	var first *Room

	// Générer 10 salles + 1 boss final
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

// Génère une salle avec ennemis en fonction de la difficulté
func generateRoom(level int) [][]string {
	size := 8
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Nombre d’ennemis croissant
	numEnemies := level + rand.Intn(3)
	for k := 0; k < numEnemies; k++ {
		for {
			x, y := rand.Intn(size), rand.Intn(size)
			if grid[x][y] == "." {
				if level == 11 {
					grid[x][y] = "👹" // Boss final
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

// --- Exploration ---
func ExploreDungeon(c *character.Character) {
	currentRoom := LoadGame()
	if currentRoom == nil {
		currentRoom = initRooms()
	}

	for {
		fmt.Printf("\n=== %s ===\n", currentRoom.Name)
		playRoom(c, currentRoom)

		if c.CurrentHP <= 0 {
			fmt.Println("💀 Vous êtes mort. Fin du jeu.")
			_ = os.Remove("save.json")
			return
		}

		if currentRoom.IsFinal && isRoomCleared(currentRoom.Grid) {
			fmt.Println("\n🎉🎉 YOU WIN! 🎉🎉")
			_ = os.Remove("save.json")
			return
		}

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
			SaveGame(currentRoom)
		} else {
			fmt.Println("❌ Direction invalide, vous restez dans la salle.")
		}
	}
}

// --- Jouer une salle ---
func fight(c *character.Character, enemy string) bool {
	enemyHP := 30
	if enemy == "👹" {
		enemyHP = 60
	}

	fmt.Printf("⚔️ Combat engagé contre %s !\n", enemy)
	for c.CurrentHP > 0 && enemyHP > 0 {
		fmt.Println("\n--- Votre tour ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser un sort")
		fmt.Println("3. Utiliser une potion")
		fmt.Println("4. Fuir")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			damage := 10
			enemyHP -= damage
			fmt.Printf("🥷 Vous attaquez et infligez %d dégâts (%d PV restants).\n", damage, enemyHP)
		case "2":
			if len(c.Skills) == 0 {
				fmt.Println("❌ Aucun sort disponible.")
			} else {
				fmt.Println("Choisissez un sort :")
				for i, s := range c.Skills {
					fmt.Printf("%d. %s\n", i+1, s)
				}
				_ = utils.AskChoice()
				// (simplifié : applique dégâts selon sort)
				damage := 15
				enemyHP -= damage
				fmt.Printf("🔥 Vous lancez un sort et infligez %d dégâts (%d PV restants).\n", damage, enemyHP)
			}
		case "3":
			character.UsePotion(c)
		case "4":
			fmt.Println("🏃 Vous fuyez le combat !")
			return false
		}

		// Vérifier si ennemi mort
		if enemyHP <= 0 {
			fmt.Printf("✅ Vous avez vaincu %s !\n", enemy)
			character.GainExp(c, 20)
			return true
		}

		// Tour de l’ennemi
		damage := 8 + (rand.Intn(5))
		c.CurrentHP -= damage
		fmt.Printf("💥 L’ennemi attaque et inflige %d dégâts (%d/%d PV).\n", damage, c.CurrentHP, c.MaxHP)
		if character.IsDead(c) {
			return false
		}
	}
	return true
}

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
			inventory.AccessInventory(c)
		case "r":
			return
		default:
			fmt.Println("Mauvais choix.")
		}

		cell := grid[playerX][playerY]
		if cell == "😈" || cell == "👹" {
			fmt.Printf("⚔️ Un ennemi %s apparaît !\n", cell)
			// choix combat simple
			fmt.Println("Voulez-vous attaquer (a) ou fuir (f) ?")
			act := utils.AskChoice()
			if act == "a" {
				fmt.Printf("🥷 Vous attaquez et éliminez %s sans prendre de dégâts.\n", cell)
				grid[playerX][playerY] = "."
				character.GainExp(c, 10)
			} else {
				fmt.Println("🏃 Vous évitez le combat, l’ennemi reste.")
			}
		}

		if isRoomCleared(grid) {
			fmt.Println("✔ Salle nettoyée !")
			return
		}
	}
}

// Affichage de la carte
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

// Sauvegarde simple
type SaveData struct {
	RoomName string `json:"room_name"`
}

func SaveGame(room *Room) {
	data := SaveData{RoomName: room.Name}
	file, err := os.Create("save.json")
	if err != nil {
		fmt.Println("⚠️ Erreur sauvegarde :", err)
		return
	}
	defer file.Close()
	_ = json.NewEncoder(file).Encode(data)
}

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

	// Reconstruire et retrouver la salle
	first := initRooms()
	queue := []*Room{first}
	seen := map[*Room]bool{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if seen[cur] {
			continue
		}
		seen[cur] = true
		if cur.Name == data.RoomName {
			return cur
		}
		for _, next := range cur.Connections {
			queue = append(queue, next)
		}
	}
	return first
}
