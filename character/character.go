package character

import (
	"fmt"
	"main/utils"
	"strings"
)

// Equipment : tête / torse / pieds
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// Character structure (exported)
type Character struct {
	Name              string
	Class             string
	Weapon            string
	Level             int
	MaxHP             int
	CurrentHP         int
	Inventory         []string
	Spell             string
	Gold              int
	Equipment         Equipment
	InventoryLimit    int
	InventoryUpgrades int
}

// InitCharacter : constructeur
func InitCharacter(name, class, weapon string, level, maxHP, currentHP int, inventory []string, spell string) Character {
	return Character{
		Name:              name,
		Class:             class,
		Weapon:            weapon,
		Level:             level,
		MaxHP:             maxHP,
		CurrentHP:         currentHP,
		Inventory:         inventory,
		Spell:             spell,
		Gold:              100,
		Equipment:         Equipment{},
		InventoryLimit:    10,
		InventoryUpgrades: 0,
	}
}

// CharacterCreation : permet à l'utilisateur de créer son perso
func CharacterCreation() Character {
	fmt.Println("\n--- Création du personnage ---")

	// Nom (lettres uniquement demandé implicitement)
	fmt.Print("Choisissez un nom : ")
	rawName := utils.AskChoice()
	name := strings.Title(strings.ToLower(rawName))

	// Classe
	fmt.Println("Choisissez une classe : (1) Humain (100 PV), (2) Elfe (80 PV), (3) Nain (120 PV)")
	classChoice := utils.AskChoice()

	class := ""
	maxHP := 100
	switch classChoice {
	case "1":
		class = "Humain"
		maxHP = 100
	case "2":
		class = "Elfe"
		maxHP = 80
	case "3":
		class = "Nain"
		maxHP = 120
	default:
		class = "Aventurier"
		maxHP = 100
	}

	// Arme de départ
	fmt.Println("Choisissez une arme : (1) Épée, (2) Arc, (3) Bâton")
	weaponChoice := utils.AskChoice()
	weapon := "Dague"
	switch weaponChoice {
	case "1":
		weapon = "Epee"
	case "2":
		weapon = "Arc"
	case "3":
		weapon = "Baton"
	}

	// PV de départ = 50% du max
	currentHP := maxHP / 2

	return InitCharacter(name, class, weapon, 1, maxHP, currentHP, []string{"Potion de Vie"}, "Coup de Poing")
}

// DisplayInfo : affiche les infos complètes du personnage
func DisplayInfo(c *Character) {
	fmt.Printf("\n--- Infos Personnage ---\n")
	fmt.Printf("Nom : %s\nClasse : %s\nArme : %s\nNiveau : %d\nPV : %d/%d\nOr : %d\nInventaire (%d/%d) : %v\nSort : %s\n",
		c.Name, c.Class, c.Weapon, c.Level, c.CurrentHP, c.MaxHP, c.Gold, len(c.Inventory), c.InventoryLimit, c.Inventory, c.Spell)

	fmt.Println("Équipement :")
	if c.Equipment.Head != "" {
		fmt.Println(" - Tête :", c.Equipment.Head)
	}
	if c.Equipment.Torso != "" {
		fmt.Println(" - Torse :", c.Equipment.Torso)
	}
	if c.Equipment.Feet != "" {
		fmt.Println(" - Pieds :", c.Equipment.Feet)
	}
}

// Equip : équipe un objet (met à jour les PV max)
func Equip(c *Character, item string) {
	switch item {
	case "Chapeau de l'aventurier":
		replaceEquipment(&c.Equipment.Head, item, c, 10)
	case "Tunique de l'aventurier":
		replaceEquipment(&c.Equipment.Torso, item, c, 25)
	case "Bottes de l'aventurier":
		replaceEquipment(&c.Equipment.Feet, item, c, 15)
	default:
		fmt.Println("❌ Cet objet ne peut pas être équipé.")
		return
	}
	fmt.Printf("✔ Vous avez équipé %s.\n", item)
}

// helper : bonus d'un équipement (par nom)
func equipmentBonus(name string) int {
	switch name {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	default:
		return 0
	}
}

// replaceEquipment : remplace un équipement dans un slot, rend l'ancien dans l'inventaire
func replaceEquipment(slot *string, newItem string, c *Character, bonus int) {
	// si slot occupé -> retirer bonus et remettre l'objet dans inventaire
	if *slot != "" {
		old := *slot
		c.MaxHP -= equipmentBonus(old)
		// remise dans inventaire (si place)
		if len(c.Inventory) < c.InventoryLimit {
			c.Inventory = append(c.Inventory, old)
		} else {
			// si pas de place, détruire (ou drop) — message
			fmt.Println("⚠️ Pas de place pour récupérer l'ancien équipement ; il est perdu.")
		}
	}
	// mettre le nouvel équipement
	*slot = newItem
	// enlever le nouvel équipement de l'inventaire (s'il y est)
	removeItemFromInventory(c, newItem)
	// appliquer bonus
	c.MaxHP += bonus
	// s'assurer que CurrentHP <= MaxHP
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
}

// removeItemFromInventory : retire 1 exemplaire d'un item si présent
func removeItemFromInventory(c *Character, item string) {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}

// Heal : soigne
func Heal(c *Character, amount int) {
	c.CurrentHP += amount
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous récupérez %d PV. (%d/%d)\n", amount, c.CurrentHP, c.MaxHP)
}

// IsDead : vérifie si mort et ressuscite à 50% du max
func IsDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Println("💀 Vous êtes mort... ressuscitation automatique (50% PV max).")
		c.CurrentHP = c.MaxHP / 2
		return true
	}
	return false
}
