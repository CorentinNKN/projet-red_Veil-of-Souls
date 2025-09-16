package inventory

import (
	"fmt"
	"main/character"
	"main/utils"
	"strconv"
	"strings"
	"time"
)

// AccessInventory : affiche inventaire et permet d'utiliser/équiper/augmenter
func AccessInventory(c *character.Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	fmt.Println("\n--- Inventaire ---")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Println("Choisissez un objet à utiliser/équiper (numéro) : ")

	choiceStr := utils.AskChoice()
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 0 || choice > len(c.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}
	if choice == 0 {
		return
	}

	item := c.Inventory[choice-1]
	lower := strings.ToLower(item)
	// decide whether to remove item from inventory after use
	removeItem := true

	switch lower {
	case "potion de vie":
		character.Heal(c, 30)
	case "potion de poison":
		poisonPot(c)
	case "augmentation d'inventaire":
		upgradeInventorySlot(c)
		// upgradeInventorySlot uses the Character state, but we must also allow removal below
	case "chapeau de l'aventurier", "tunique de l'aventurier", "bottes de l'aventurier":
		// Equip handles removing the item from inventory internally
		character.Equip(c, item)
		removeItem = false
	default:
		fmt.Println("Cet objet n'a aucun effet utilisable pour le moment.")
	}

	if removeItem {
		// retirer l'objet utilisé
		if choice-1 >= 0 && choice-1 < len(c.Inventory) {
			c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)
		}
	}
}

// poisonPot : inflige 10 PV/s pendant 3s, affiche PV à chaque tick
func poisonPot(c *character.Character) {
	fmt.Println("⚗️ Vous avez bu une Potion de Poison !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf("☠️ Seconde %d : %d/%d PV\n", i, c.CurrentHP, c.MaxHP)
	}
	if c.CurrentHP <= 0 {
		character.IsDead(c)
	}
}

// upgradeInventorySlot : augmente la capacité max de l'inventaire (+10) (max 3 fois)
func upgradeInventorySlot(c *character.Character) {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("❌ Vous ne pouvez plus augmenter l'inventaire.")
		return
	}
	c.InventoryLimit += 10
	c.InventoryUpgrades++
	fmt.Printf("✔ Inventaire augmenté ! Nouvelle capacité : %d\n", c.InventoryLimit)
}
