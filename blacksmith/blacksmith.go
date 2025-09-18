package blacksmith

import (
	"fmt"
	"main/character"
	"main/utils"
)

// AccessBlacksmith : affiche le menu du forgeron et permet au joueur d'améliorer son stuff
func AccessBlacksmith(c *character.Character) {
	for {
		// Affichage du menu principal du forgeron
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("💰 Or disponible : %d\n", c.Gold)
		fmt.Println("1. Améliorer l’inventaire (+10 slots, max 3 fois) (50 or)")
		fmt.Println("2. Forger un casque (+2 PV max) (30 or)")
		fmt.Println("3. Forger une armure (+5 PV max) (50 or)")
		fmt.Println("4. Forger des bottes (+1 PV max) (20 or)")
		fmt.Println("5. Quitter")

		// Lecture du choix de l'utilisateur
		choice := utils.AskChoice()

		switch choice {
		// Amélioration de l’inventaire
		case "1":
			if c.Gold >= 50 {
				// Appelle la fonction UpgradeInventory (limité à 3 fois)
				if character.UpgradeInventory(c) {
					c.Gold -= 50
				}
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}

		// Forger un casque
		case "2":
			if c.Gold >= 30 {
				// Ajoute un casque et augmente les PV max du perso
				c.Equipment.Head = "Casque"
				c.MaxHP += 2
				c.CurrentHP += 2
				c.Gold -= 30
				fmt.Println("🪓 Vous équipez un casque (+2 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}

		// Forger une armure
		case "3":
			if c.Gold >= 50 {
				// Ajoute une armure et booste les PV max
				c.Equipment.Torso = "Armure"
				c.MaxHP += 5
				c.CurrentHP += 5
				c.Gold -= 50
				fmt.Println("🛡️ Vous équipez une armure (+5 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}

		// Forger des bottes
		case "4":
			if c.Gold >= 20 {
				// Ajoute des bottes et augmente légèrement les PV
				c.Equipment.Feet = "Bottes"
				c.MaxHP += 1
				c.CurrentHP += 1
				c.Gold -= 20
				fmt.Println("🥾 Vous équipez des bottes (+1 PV max).")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}

		// Quitter le menu
		case "5":
			fmt.Println("Vous quittez le forgeron.")
			return

		// Gestion des entrées invalides
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
