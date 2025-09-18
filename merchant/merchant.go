package merchant

import (
	"fmt"
	"main/character"
	"main/utils"
)

// AccessMerchant : menu marchand
func AccessMerchant(c *character.Character) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("💰 Or disponible : %d\n", c.Gold)
		fmt.Println("1. Acheter une potion de vie (20 or)")
		fmt.Println("2. Acheter une potion de poison (15 or)")
		fmt.Println("3. Vendre un objet (10 or)")
		fmt.Println("4. Quitter")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			if c.Gold >= 20 {
				if character.AddItem(c, "Potion de vie") {
					c.Gold -= 20
				}
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "2":
			if c.Gold >= 15 {
				if character.AddItem(c, "Potion de poison") {
					c.Gold -= 15
				}
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "3":
			if len(c.Inventory) == 0 {
				fmt.Println("❌ Inventaire vide, rien à vendre.")
				continue
			}
			fmt.Println("Votre inventaire :", c.Inventory)
			fmt.Print("Quel objet voulez-vous vendre ? ")
			item := utils.AskChoice()
			if character.RemoveItem(c, item) {
				c.Gold += 10
				fmt.Printf("💰 Vous vendez %s et gagnez 10 or.\n", item)
			} else {
				fmt.Println("❌ Objet introuvable.")
			}
		case "4":
			fmt.Println("Vous quittez le marchand.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
