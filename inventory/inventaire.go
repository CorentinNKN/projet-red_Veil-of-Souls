package inventory

import (
	"fmt"
	"main/character"
	"main/utils"
)

func AccessInventory(c *character.Character) {
	for {
		fmt.Println("\n--- Inventaire ---")
		if len(c.Inventory) == 0 {
			fmt.Println("Inventaire vide.")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}
		fmt.Println("a. Utiliser une potion")
		fmt.Println("b. Retour")

		choice := utils.AskChoice()

		switch choice {
		case "a":
			TakePot(c)
		case "b":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func TakePot(c *character.Character) {
	index := -1
	for i, item := range c.Inventory {
		if item == "Potion" {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Aucune potion disponible.")
		return
	}

	// Retirer la potion
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	// Soigner le personnage
	character.Heal(c, 50)
}
