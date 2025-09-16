package combat

import (
	"fmt"
	"main/character"
	"main/inventory"
	"main/monster"
	"main/utils"
	"math/rand"
)

// TrainingCombat : combat tour-par-tour contre un gobelin
func TrainingCombat(c *character.Character) {
	m := monster.InitGoblin()
	fmt.Println("\n--- Combat d'entraînement : Gobelin ---")
	for m.CurrentHP > 0 && c.CurrentHP > 0 {
		// états
		fmt.Printf("\n%s : %d/%d PV | %s : %d/%d PV\n", c.Name, c.CurrentHP, c.MaxHP, m.Name, m.CurrentHP, m.MaxHP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser objet")
		fmt.Println("3. Fuir")

		choice := utils.AskChoice()
		switch choice {
		case "1":
			// calcul dégâts joueur
			base := rand.Intn(6) + 5
			weaponBonus := weaponBonus(c.Weapon)
			damage := base + weaponBonus + c.Level*2
			m.CurrentHP -= damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
			fmt.Printf("Vous infligez %d dégâts au %s (%d/%d).\n", damage, m.Name, m.CurrentHP, m.MaxHP)
		case "2":
			// ouvrir inventaire et permettre d'utiliser potions
			inventory.AccessInventory(c)
		case "3":
			fmt.Println("Vous fuyez le combat.")
			return
		default:
			fmt.Println("Choix invalide.")
		}

		// si monstre mort
		if m.CurrentHP <= 0 {
			fmt.Println("✔ Vous avez vaincu le gobelin d'entraînement !")
			// récompense : or + ressources éventuelles
			rewardGold := 10
			c.Gold += rewardGold
			fmt.Printf("Vous gagnez %d or.\n", rewardGold)
			// loot aléatoire
			if rand.Intn(100) < 50 { // 50% chance
				if len(c.Inventory) < c.InventoryLimit {
					c.Inventory = append(c.Inventory, "Fourrure de Loup")
					fmt.Println("Vous récupérez : Fourrure de Loup.")
				} else {
					fmt.Println("Votre inventaire est plein, le loot est perdu.")
				}
			}
			return
		}

		// tour du monstre
		fmt.Printf("%s attaque !\n", m.Name)
		c.CurrentHP -= m.Attack
		fmt.Printf("Vous subissez %d PV de dégâts (%d/%d).\n", m.Attack, c.CurrentHP, c.MaxHP)
		if character.IsDead(c) {
			fmt.Println("⚡ Vous avez été ressuscité et le combat s'arrête.")
			return
		}
	}
}

// weaponBonus : simple bonus selon arme
func weaponBonus(weapon string) int {
	switch weapon {
	case "Epee":
		return 5
	case "Arc":
		return 3
	case "Baton":
		return 2
	default:
		return 0
	}
}
