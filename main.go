package main

import (
	"fmt"
	"veil_of_souls/character"
	"veil_of_souls/intro"
	"veil_of_souls/inventory"
	"veil_of_souls/merchant"
	"veil_of_souls/utils"
)

func main() {
	// Page d'intro
	intro.ShowIntro()

	// Création du personnage
	c1 := character.InitCharacter("TonNom", "Elfe", 1, 100, 40, []string{"Potion de vie", "Potion de vie", "Potion de vie"})

	// Boucle du menu principal
	for {
		fmt.Println("\n----- Menu Principal -----")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			character.DisplayInfo(&c1)
		case 2:
			inventory.AccessInventory(&c1)
		case 3:
			merchant.Marchand(&c1)
		case 4:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}

		utils.IsDead(&c1)
	}
}
