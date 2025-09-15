package merchant

import (
	"fmt"
	"main/character"
	"main/utils"
)

func AccessMerchant(c *character.Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Acheter une potion de vie (gratuit)")
		fmt.Println("2. Retour")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			c.Inventory = append(c.Inventory, "Potion")
			fmt.Println("Vous avez obtenu une potion !")
		case "2":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
