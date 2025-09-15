package main

func main() {
	character := initCharacter("Aragorn", "Guerrier", 5, 100, 75, []string{"Potion de soin", "Épée", "Bouclier"})
	DisplayInfo(&character)
	accesInventory(&character)
}
