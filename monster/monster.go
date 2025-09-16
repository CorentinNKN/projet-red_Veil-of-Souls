package monster

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin d’entraînement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}
