package game

import (
	"fmt"
	"main/blacksmith"
	"main/character"
	"main/combat"
	"main/inventory"
	"main/mapgame"
	"main/merchant"
	"main/utils"
	"os"
)

func StartGame() {
	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1. Nouvelle Partie")
		fmt.Println("2. Charger la sauvegarde")
		fmt.Println("3. Quitter")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			player := character.CharacterCreation()
			mainMenu(&player)
		case "2":
			if _, err := os.Stat("save.json"); os.IsNotExist(err) {
				fmt.Println("❌ Aucune sauvegarde trouvée.")
			} else {
				player := character.CharacterCreation() // ⚠️ tu peux plus tard sauver le joueur aussi
				mapgame.ExploreDungeon(&player)
			}
		case "3":
			fmt.Println("Merci d'avoir joué à Veil of Souls !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func mainMenu(player *character.Character) {
	for {
		fmt.Println("\n--- Menu du Joueur ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Accéder au forgeron")
		fmt.Println("5. Explorer le donjon")
		fmt.Println("6. Combat d'entraînement")
		fmt.Println("7. Retour au menu principal")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			character.DisplayInfo(player)
		case "2":
			inventory.AccessInventory(player)
		case "3":
			merchant.AccessMerchant(player)
		case "4":
			blacksmith.AccessBlacksmith(player)
		case "5":
			mapgame.ExploreDungeon(player)
		case "6":
			combat.TrainingCombat(player)
		case "7":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
