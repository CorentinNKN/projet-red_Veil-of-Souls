package merchant

import (
	"fmt"
	"main/character"
	"main/utils"
)

type SaleItem struct {
	Name string
	Cost int
}

var itemsForSale = []SaleItem{
	{"Potion de Vie", 3},
	{"Potion de Poison", 6},
	{"Livre de Sort : Boule de feu", 25},
	{"Fourrure de Loup", 4},
	{"Peau de Troll", 7},
	{"Cuir de Sanglier", 3},
	{"Plume de Corbeau", 1},
	{"Augmentation d'inventaire", 30},
}

func AccessMerchant(c *character.Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Printf("💰 Or disponible : %d pièces\n", c.Gold)
		for i, it := range itemsForSale {
			fmt.Printf("%d. %s (%d or)\n", i+1, it.Name, it.Cost)
		}
		fmt.Printf("%d. Quitter le marchand\n", len(itemsForSale)+1)

		choice := utils.AskChoiceIndex(len(itemsForSale) + 1)
		if choice == len(itemsForSale)+1 {
			return
		}

		item := itemsForSale[choice-1]
		if c.Gold < item.Cost {
			fmt.Println("❌ Pas assez d'or pour cet achat.")
			continue
		}
		if len(c.Inventory) >= c.InventoryLimit {
			fmt.Println("❌ Inventaire plein, impossible d'acheter.")
			continue
		}

		// Achat
		c.Gold -= item.Cost
		c.Inventory = append(c.Inventory, item.Name)
		fmt.Printf("✔ Vous avez acheté %s pour %d or. Or restant : %d\n", item.Name, item.Cost, c.Gold)
	}
}
