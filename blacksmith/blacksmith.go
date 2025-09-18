package blacksmith

import (
	"fmt"
	"main/character"
	"main/utils"
)

// AccessBlacksmith permet d'accéder au forgeron
func AccessBlacksmith(c *character.Character) {
	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("Or disponible : %d\n", c.Gold)
		fmt.Println("Équipement actuel :")
		fmt.Printf("Tête : %s | Torse : %s | Pieds : %s\n", c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet)
		fmt.Println("\nOptions de craft :")
		fmt.Println("1. Chapeau (+10 PV max) - 10 or + 2 Plumes")
		fmt.Println("2. Tunique (+25 PV max) - 20 or + 2 Fourrures")
		fmt.Println("3. Bottes (+15 PV max) - 15 or + 2 Cuirs")
		fmt.Println("4. Quitter le forgeron")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			craftItem(c, "Chapeau", "Plume", 2, 10, 10, &c.Equipment.Head)
		case "2":
			craftItem(c, "Tunique", "Fourrure", 2, 20, 25, &c.Equipment.Torso)
		case "3":
			craftItem(c, "Bottes", "Cuir", 2, 15, 15, &c.Equipment.Feet)
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// craftItem permet de fabriquer un équipement
func craftItem(c *character.Character, itemName, resource string, resNeeded, cost, hpBoost int, slot *string) {
	if *slot != "" {
		fmt.Printf("❌ Vous portez déjà un %s.\n", itemName)
		return
	}
	if c.Gold < cost {
		fmt.Println("❌ Pas assez d’or.")
		return
	}

	// Vérifie si le joueur a assez de ressources
	count := 0
	for _, it := range c.Inventory {
		if it == resource {
			count++
		}
	}
	if count < resNeeded {
		fmt.Printf("❌ Vous avez besoin de %d %s (vous en avez %d).\n", resNeeded, resource, count)
		return
	}

	// Retire les ressources
	removed := 0
	newInv := []string{}
	for _, it := range c.Inventory {
		if it == resource && removed < resNeeded {
			removed++
			continue
		}
		newInv = append(newInv, it)
	}
	c.Inventory = newInv

	// Retire l’or
	c.Gold -= cost

	// Ajoute l’équipement
	*slot = itemName
	c.MaxHP += hpBoost
	c.CurrentHP += hpBoost

	fmt.Printf("✅ Vous avez fabriqué un %s (+%d PV max).\n", itemName, hpBoost)
}
