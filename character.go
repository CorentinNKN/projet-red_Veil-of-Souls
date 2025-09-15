package main

import (
	"fmt"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

// Fonction pour utiliser une potion de soin
func takePot(c *Character, index int) {
	healAmount := 20
	if c.PVActuels < c.PVMax {
		c.PVActuels += healAmount
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Printf("Vous avez utilisé une Potion de soin. PV actuels : %d/%d\n", c.PVActuels, c.PVMax)
		// Retirer la potion de l'inventaire
		c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
	} else {
		fmt.Println("Vos PV sont déjà au maximum.")
	}
}

func initCharacter(nom string, classe string, niveau int, PvMax int, PvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      PvMax,
		PVActuels:  PvActuels,
		Inventaire: inventaire,
	}
}

func DisplayInfo(c *Character) {
	fmt.Printf("Nom: %s\n", c.Nom)
	fmt.Printf("Classe: %s\n", c.Classe)
	fmt.Printf("Niveau: %d\n", c.Niveau)
	fmt.Printf("PV: %d/%d\n", c.PVActuels, c.PVMax)
	fmt.Println("--------------------------------------")
}

// affichage de l'inventaire

func accesInventory(c *Character) {
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("-----Inventaire-----")
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("--------------------")

	var choix int
	fmt.Println("Voulez-vous utiliser une potion ? (tapez le numéro 0 pour retour)")
	fmt.Scan(&choix)

	if choix > 0 && choix <= len(c.Inventaire) {
		if c.Inventaire[choix-1] == "Potion de soin" {
			takePot(c, choix-1)
		} else {
			fmt.Println("Cet objet n'ai pas utilisable.")
		}
	}
}
