package merchant

import (
	"fmt"
	"main/character"
	"main/utils"
)

// AccessMerchant ouvre la boutique
func AccessMerchant(c *character.Character) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Or disponible : %d\n", c.Gold)
		fmt.Println("Articles en vente :")
		fmt.Println("1. Potion de vie (3 or)")
		fmt.Println("2. Potion de poison (6 or)")
		fmt.Println("3. Livre de sort : Boule de feu (25 or)")
		fmt.Println("4. Ressources (Fourrure, Cuir, Plume...) (5 or)")
		fmt.Println("5. Augmentation inventaire (+10 slots, max 3) (30 or)")
		fmt.Println("6. Quitter le marchand")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			buyItem(c, "Potion de vie", 3)
		case "2":
			buyItem(c, "Potion de poison", 6)
		case "3":
			if c.Gold < 25 {
				fmt.Println("❌ Pas assez d’or.")
				continue
			}
			if character.LearnSpell(c, "Boule de feu") {
				c.Gold -= 25
			}
		case "4":
			fmt.Println("Quelle ressource voulez-vous acheter ? (Fourrure, Cuir, Plume...)")
			res := utils.AskChoice()
			buyItem(c, res, 5)
		case "5":
			if c.Gold < 30 {
				fmt.Println("❌ Pas assez d’or.")
				continue
			}
			if character.UpgradeInventory(c) {
				c.Gold -= 30
			}
		case "6":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// --- Aides internes ---
func buyItem(c *character.Character, item string, cost int) {
	if c.Gold < cost {
		fmt.Println("❌ Pas assez d’or.")
		return
	}
	if character.AddItem(c, item) {
		c.Gold -= cost
	}
}
