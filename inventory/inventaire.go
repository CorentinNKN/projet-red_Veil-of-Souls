package inventory

import (
	"fmt"
	"main/character"
	"main/utils"
	"strings"
)

// AccessInventory permet au joueur de voir et gérer son inventaire
func AccessInventory(c *character.Character) {
	for {
		fmt.Println("\n=== Inventaire ===")
		fmt.Printf("Contenu (%d/%d) : %v\n", len(c.Inventory), c.InventoryCapacity, c.Inventory)
		fmt.Println("Options :")
		fmt.Println("1. Utiliser une potion de vie (+50 PV)")
		fmt.Println("2. Utiliser une potion de poison (10 dégâts/s pendant 3s)")
		fmt.Println("3. Supprimer un objet")
		fmt.Println("4. Utiliser un sort")
		fmt.Println("5. Quitter l’inventaire")

		choice := utils.AskChoice()

		switch choice {
		case "1":
			character.UsePotion(c)
		case "2":
			character.UsePoisonPot(c)
		case "3":
			fmt.Print("Nom de l’objet à supprimer : ")
			item := utils.AskChoice()
			if !character.RemoveItem(c, item) {
				fmt.Println("❌ Objet introuvable.")
			} else {
				fmt.Printf("🗑️ %s supprimé de l’inventaire.\n", item)
			}
		case "4":
			useSpell(c)
		case "5":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// --- Sorts / compétences ---
func useSpell(c *character.Character) {
	if len(c.Skills) == 0 {
		fmt.Println("❌ Vous n’avez appris aucun sort.")
		return
	}

	fmt.Println("\n=== Sorts disponibles ===")
	for i, s := range c.Skills {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	fmt.Print("Choisissez un sort : ")
	choice := utils.AskChoice()

	idx := -1
	for i, s := range c.Skills {
		if choice == fmt.Sprint(i+1) || strings.EqualFold(choice, s) {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("❌ Choix invalide.")
		return
	}

	spell := c.Skills[idx]
	switch spell {
	case "Coup de poing":
		fmt.Println("👊 Vous donnez un coup de poing (8 dégâts).")
	case "Boule de feu":
		if c.CurrentMana < 10 {
			fmt.Println("❌ Pas assez de mana (10 requis).")
			return
		}
		c.CurrentMana -= 10
		fmt.Println("🔥 Vous lancez une boule de feu (18 dégâts). Mana -10.")
	default:
		fmt.Printf("✨ Vous utilisez le sort : %s\n", spell)
	}
}
