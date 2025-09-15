package character

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func DisplayInfo(c *Character) {
	fmt.Printf("\n--- Infos Personnage ---\n")
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nInventaire: %v\n",
		c.Name, c.Class, c.Level, c.CurrentHP, c.MaxHP, c.Inventory)
}

func Heal(c *Character, amount int) {
	c.CurrentHP += amount
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("%s a récupéré de la vie (%d/%d PV).\n", c.Name, c.CurrentHP, c.MaxHP)
}

func IsDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Printf("%s est mort... mais ressuscite avec 50%% de ses PV max !\n", c.Name)
		c.CurrentHP = c.MaxHP / 2
		return true
	}
	return false
}
