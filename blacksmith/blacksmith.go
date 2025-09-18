package blacksmith

import (
	"fmt"
	"main/character"
	"main/utils"
)

// AccessBlacksmith : menu forgeron
func AccessBlacksmith(c *character.Character) {
	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("💰 Or disponible : %d\n", c.Gold)
		fmt.Println("1. Améliorer l’inventaire (+10 slots, max 3 fois) (50 or)")
		fmt.Println("2. Forger un casque (+2 PV max) (30 or)")
		fmt.Println("3. Forger une armure (+5 PV max) (50 or)")
		fmt.Println("4. Forger des bottes (+1 PV max) (20 or)")
		fmt.Println("5. Quitter")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			if c.Gold >= 50 {
				if character.UpgradeInventory(c) {
					c.Gold -= 50
				}
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "2":
			if c.Gold >= 30 {
				c.Equipment.Head = "Casque"
				c.MaxHP += 2
				c.CurrentHP += 2
				c.Gold -= 30
				fmt.Println("🪓 Vous équipez un casque (+2 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "3":
			if c.Gold >= 50 {
				c.Equipment.Torso = "Armure"
				c.MaxHP += 5
				c.CurrentHP += 5
				c.Gold -= 50
				fmt.Println("🛡️ Vous équipez une armure (+5 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "4":
			if c.Gold >= 20 {
				c.Equipment.Feet = "Bottes"
				c.MaxHP += 1
				c.CurrentHP += 1
				c.Gold -= 20
				fmt.Println("🥾 Vous équipez des bottes (+1 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case "5":
			fmt.Println("Vous quittez le forgeron.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
