package blacksmith

import (
	"fmt"
	"main/character"
	"main/utils"
)

func AccessBlacksmith(c *character.Character) {
	for {
		fmt.Println("\n--- Forgeron ---")
		fmt.Printf("💰 Or disponible : %d pièces\n", c.Gold)
		fmt.Println("1. Fabriquer Chapeau de l'aventurier (5 or)")
		fmt.Println("2. Fabriquer Tunique de l'aventurier (5 or)")
		fmt.Println("3. Fabriquer Bottes de l'aventurier (5 or)")
		fmt.Println("4. Retour")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			craftItem(c, "Chapeau de l'aventurier", []string{"Plume de Corbeau", "Cuir de Sanglier"})
		case "2":
			craftItem(c, "Tunique de l'aventurier", []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"})
		case "3":
			craftItem(c, "Bottes de l'aventurier", []string{"Fourrure de Loup", "Cuir de Sanglier"})
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func craftItem(c *character.Character, item string, requirements []string) {
	// coût en or : 5
	if c.Gold < 5 {
		fmt.Println("❌ Pas assez d'or pour fabriquer cet objet.")
		return
	}
	if len(c.Inventory) >= c.InventoryLimit {
		fmt.Println("❌ Inventaire plein.")
		return
	}
	// vérifier ressources
	for _, req := range requirements {
		if !hasItem(c, req) {
			fmt.Printf("❌ Il vous manque %s pour fabriquer %s.\n", req, item)
			return
		}
	}
	// consommer ressources
	for _, req := range requirements {
		removeItem(c, req)
	}
	// retirer or
	c.Gold -= 5
	// ajouter l'équipement
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✔ Vous avez fabriqué %s ! (-5 or)\n", item)
}

func hasItem(c *character.Character, item string) bool {
	for _, v := range c.Inventory {
		if v == item {
			return true
		}
	}
	return false
}

func removeItem(c *character.Character, item string) {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}
