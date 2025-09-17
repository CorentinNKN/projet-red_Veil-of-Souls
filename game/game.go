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
)

func StartGame() {
	player := character.CharacterCreation()

	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Accéder au forgeron")
		fmt.Println("5. Explorer le donjon")
		fmt.Println("6. Combat d'entraînement")
		fmt.Println("7. Quitter")

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
			mapgame.ExploreDungeon(&player) // ✅ correction ici
		case "6":
			combat.TrainingCombat(&player)
		case "7":
			fmt.Println("Merci d'avoir joué à Veil of Souls !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
