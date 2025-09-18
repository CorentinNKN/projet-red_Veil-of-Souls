package character

import (
	"fmt"
	"main/utils"
	"strings"
	"time"
	"unicode"
)

// Equipment : slots d’équipement du joueur
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// Character : toutes les infos du personnage
type Character struct {
	Name                 string
	Class                string
	Level                int
	MaxHP                int
	CurrentHP            int
	MaxMana              int
	CurrentMana          int
	Inventory            []string
	InventoryCapacity    int
	InventoryUpgradeUses int // nb max d’augmentations d’inventaire (3 fois)
	Gold                 int
	Skills               []string
	Equipment            Equipment
	Exp                  int
	ExpMax               int
	Initiative           int
}

// Création du personnage
func CharacterCreation() Character {
	fmt.Println("=== Création du personnage ===")
	var name string

	// Demande un nom valide (lettres uniquement)
	for {
		fmt.Print("Entrez le nom du personnage (lettres uniquement) : ")
		input := utils.AskChoice()
		input = strings.TrimSpace(input)
		if validName(input) {
			name = normalizeName(input) // Nom propre (ex: "Jean Dupont")
			break
		}
		fmt.Println("Nom invalide : n'utilisez que des lettres.")
	}

	// Choix de la classe
	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Humain (100 PV max)")
	fmt.Println("2. Elfe  (80 PV max)")
	fmt.Println("3. Nain  (120 PV max)")

	var class string
	var maxHP int
	for {
		choice := utils.AskChoice()
		switch choice {
		case "1", "Humain", "humain", "HUMAIN":
			class = "Humain"
			maxHP = 100
		case "2", "Elfe", "elfe", "ELFE":
			class = "Elfe"
			maxHP = 80
		case "3", "Nain", "nain", "NAIN":
			class = "Nain"
			maxHP = 120
		default:
			fmt.Println("Choix invalide, réessayez.")
			continue
		}
		break
	}

	// Statistiques de base
	level := 1
	currentHP := maxHP / 2 // spawn à 50% PV
	maxMana := maxHP / 4
	if maxMana < 10 {
		maxMana = 10
	}
	currentMana := maxMana

	// Inventaire par défaut
	inventory := []string{"Potion de vie", "Potion de vie", "Potion de vie"}

	// Expérience et or
	exp := 0
	expMax := 100
	gold := 100

	// Sort de base
	skills := []string{"Coup de poing"}
	initiative := 10

	// Création du personnage
	c := Character{
		Name:                 name,
		Class:                class,
		Level:                level,
		MaxHP:                maxHP,
		CurrentHP:            currentHP,
		MaxMana:              maxMana,
		CurrentMana:          currentMana,
		Inventory:            inventory,
		InventoryCapacity:    10,
		InventoryUpgradeUses: 0,
		Gold:                 gold,
		Skills:               skills,
		Equipment:            Equipment{},
		Exp:                  exp,
		ExpMax:               expMax,
		Initiative:           initiative,
	}

	fmt.Printf("\nPersonnage créé : %s (%s) - Niveau %d\n", c.Name, c.Class, c.Level)
	fmt.Printf("PV : %d/%d | Mana : %d/%d | Or : %d\n", c.CurrentHP, c.MaxHP, c.CurrentMana, c.MaxMana, c.Gold)
	return c
}

// Affichage
func DisplayInfo(c *Character) {
	fmt.Println("\n=== Informations du personnage ===")
	fmt.Printf("Nom : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Mana : %d / %d\n", c.CurrentMana, c.MaxMana)
	fmt.Printf("Exp : %d / %d\n", c.Exp, c.ExpMax)
	fmt.Printf("Or : %d\n", c.Gold)
	fmt.Printf("Inventaire (%d/%d) : %v\n", len(c.Inventory), c.InventoryCapacity, c.Inventory)
	fmt.Printf("Equipement : Tête=%s Torse=%s Pieds=%s\n",
		emptyOr(c.Equipment.Head), emptyOr(c.Equipment.Torso), emptyOr(c.Equipment.Feet))
	fmt.Printf("Sorts appris : %v\n", c.Skills)
	fmt.Printf("Initiative : %d\n", c.Initiative)
}

// Mort et résurrection
func IsDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = c.MaxHP / 2
		if c.CurrentMana < c.MaxMana/2 {
			c.CurrentMana = c.MaxMana / 2
		}
		fmt.Printf("💀 %s est mort, mais est ressuscité à %d/%d PV.\n", c.Name, c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}

// Potions
func UsePotion(c *Character) bool {
	for i, item := range c.Inventory {
		if strings.ToLower(item) == "potion de vie" || strings.ToLower(item) == "potion" {
			heal := 50
			c.CurrentHP += heal
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("🍷 Vous buvez une potion et regagnez %d PV (%d/%d).\n", heal, c.CurrentHP, c.MaxHP)
			return true
		}
	}
	fmt.Println("❌ Aucune potion de vie dans l'inventaire.")
	return false
}

func UsePoisonPot(c *Character) bool {
	for i, item := range c.Inventory {
		if strings.ToLower(item) == "potion de poison" || strings.ToLower(item) == "poison" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Println("☠️ Vous avez utilisé une potion de poison (10 dégâts/s pendant 3s).")
			for t := 1; t <= 3; t++ {
				c.CurrentHP -= 10
				if c.CurrentHP < 0 {
					c.CurrentHP = 0
				}
				fmt.Printf("Dégâts de poison %ds : PV %d/%d\n", t, c.CurrentHP, c.MaxHP)
				time.Sleep(1 * time.Second)
				if IsDead(c) {
					break
				}
			}
			return true
		}
	}
	fmt.Println("❌ Aucune potion de poison dans l'inventaire.")
	return false
}

// Helpers
func validName(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' {
			return false
		}
	}
	return true
}

func normalizeName(s string) string {
	s = strings.TrimSpace(s)
	parts := strings.Fields(s)
	for i, p := range parts {
		runes := []rune(strings.ToLower(p))
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}
	return strings.Join(parts, " ")
}

func emptyOr(s string) string {
	if s == "" {
		return "aucun"
	}
	return s
}

// Inventaire
func RemoveItem(c *Character, item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

// Progression
func GainExp(c *Character, exp int) {
	c.Exp += exp
	fmt.Printf("🔹 Vous gagnez %d XP (%d/%d)\n", exp, c.Exp, c.Level*100)

	// Vérifie si on monte de niveau
	for c.Exp >= c.Level*100 {
		c.Exp -= c.Level * 100
		c.Level++
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		c.MaxMana += 5
		c.CurrentMana = c.MaxMana
		fmt.Printf("🎉 Niveau %d atteint ! PV=%d Mana=%d\n", c.Level, c.MaxHP, c.MaxMana)

		// Débloque des sorts selon le niveau
		switch c.Level {
		case 3:
			LearnSpell(c, "Boule de feu")
		case 5:
			LearnSpell(c, "Éclair")
		case 8:
			LearnSpell(c, "Explosion")
		}
	}
}

// Améliorations
func UpgradeInventory(c *Character) bool {
	if c.InventoryUpgradeUses >= 3 {
		return false
	}
	c.InventoryCapacity += 10
	c.InventoryUpgradeUses++
	return true
}

func AddItem(c *Character, item string) bool {
	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Println("❌ Inventaire plein.")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✅ %s ajouté à l'inventaire.\n", item)
	return true
}

func LearnSpell(c *Character, spell string) bool {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Println("❌ Sort déjà connu.")
			return false
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Printf("✅ Sort %s appris.\n", spell)
	return true
}
