package combat

import (
	"fmt"
	"main/character"
	"main/inventory"
	"main/utils"
	"math/rand"
)

// TrainingCombat lance un combat d'entraînement contre un gobelin
func TrainingCombat(c *character.Character) {
	fmt.Println("\n=== Combat d'entraînement ===")
	goblin := Enemy{
		Name:       "Gobelin",
		HP:         50,
		Attack:     10,
		CritTurn:   3, // tous les 3 tours = critique
		ExpReward:  20,
		GoldReward: 10,
	}

	runBattle(c, &goblin)
}

// Enemy représente un adversaire
type Enemy struct {
	Name       string
	HP         int
	Attack     int
	CritTurn   int
	ExpReward  int
	GoldReward int
}

// Combat générique contre un ennemi
func runBattle(c *character.Character, e *Enemy) {
	turn := 1

	// initiative aléatoire : joueur ou ennemi commence
	playerTurn := rand.Intn(2) == 0
	if playerTurn {
		fmt.Println("👉 Vous avez l'initiative.")
	} else {
		fmt.Println("👹 L'ennemi attaque en premier !")
	}

	for c.CurrentHP > 0 && e.HP > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)

		if playerTurn {
			playerAction(c, e)
		} else {
			enemyAction(c, e, turn)
		}

		// check morts
		if e.HP <= 0 {
			fmt.Printf("🏆 Vous avez vaincu %s !\n", e.Name)
			c.Gold += e.GoldReward
			character.GainExp(c, e.ExpReward)
			fmt.Printf("Récompenses : %d or, %d exp.\n", e.GoldReward, e.ExpReward)
			return
		}
		if c.CurrentHP <= 0 {
			if character.IsDead(c) {
				fmt.Println("⚡ Vous êtes ressuscité et pouvez continuer le combat.")
			} else {
				fmt.Println("💀 Vous êtes mort.")
				return
			}
		}

		// switch de tour
		playerTurn = !playerTurn
		turn++
	}
}

// Actions possibles du joueur
func playerAction(c *character.Character, e *Enemy) {
	fmt.Printf("\nVos PV : %d/%d | Mana : %d/%d\n", c.CurrentHP, c.MaxHP, c.CurrentMana, c.MaxMana)
	fmt.Printf("PV %s : %d\n", e.Name, e.HP)
	fmt.Println("Actions disponibles :")
	fmt.Println("1. Attaquer (Coup de poing 8 dmg)")
	fmt.Println("2. Sorts")
	fmt.Println("3. Inventaire")
	fmt.Println("4. Fuir")

	choice := utils.AskChoice()
	switch choice {
	case "1":
		fmt.Println("👊 Vous attaquez avec un coup de poing !")
		e.HP -= 8
	case "2":
		spellChoice(c, e)
	case "3":
		inventory.AccessInventory(c)
	case "4":
		fmt.Println("🏃 Vous fuyez le combat.")
		e.HP = 0
	default:
		fmt.Println("❌ Choix invalide, vous perdez votre tour.")
	}
}

// Sorts disponibles
func spellChoice(c *character.Character, e *Enemy) {
	if len(c.Skills) == 0 {
		fmt.Println("❌ Aucun sort appris.")
		return
	}

	fmt.Println("\n--- Sorts ---")
	for i, s := range c.Skills {
		fmt.Printf("%d. %s\n", i+1, s)
	}
	choice := utils.AskChoice()

	switch choice {
	case "1", "Coup de poing":
		fmt.Println("👊 Coup de poing ! (8 dmg)")
		e.HP -= 8
	case "2", "Boule de feu":
		if !hasSkill(c, "Boule de feu") {
			fmt.Println("❌ Vous n’avez pas appris ce sort.")
			return
		}
		if c.CurrentMana < 10 {
			fmt.Println("❌ Pas assez de mana (10 requis).")
			return
		}
		c.CurrentMana -= 10
		fmt.Println("🔥 Boule de feu lancée ! (18 dmg)")
		e.HP -= 18
	default:
		fmt.Println("❌ Sort invalide.")
	}
}

// Vérifie si un skill est appris
func hasSkill(c *character.Character, skill string) bool {
	for _, s := range c.Skills {
		if s == skill {
			return true
		}
	}
	return false
}

// Action de l'ennemi
func enemyAction(c *character.Character, e *Enemy, turn int) {
	dmg := e.Attack
	if e.CritTurn > 0 && turn%e.CritTurn == 0 {
		dmg *= 2
		fmt.Printf("💥 %s fait une attaque critique !\n", e.Name)
	}
	c.CurrentHP -= dmg
	fmt.Printf("👹 %s vous inflige %d dégâts (%d/%d PV).\n", e.Name, dmg, c.CurrentHP, c.MaxHP)
}
