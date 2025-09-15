package main

import (
	"fmt"
	"main/character"
	"main/intro"
	"main/inventory"
	"main/mapgame"
	"main/merchant"
	"main/utils"
)

func main() {
	intro.ShowIntro()

	// Initialisation du personnage
	c1 := character.InitCharacter("TonNom", "Elfe", 1, 100, 40, []string{"Potion", "Potion", "Potion"})

	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Explorer la carte")
		fmt.Println("5. Quitter")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			character.DisplayInfo(&c1)
		case "2":
			inventory.AccessInventory(&c1)
		case "3":
			merchant.AccessMerchant(&c1)
		case "4":
			mapgame.StartExploration(&c1)
		case "5":
			fmt.Println("Merci d’avoir joué à Veil of Souls !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}
