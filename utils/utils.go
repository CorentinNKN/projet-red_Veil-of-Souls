package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Lit une ligne depuis stdin et la trim
func AskChoice() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Lit un choix numérique entre 1 et max (boucle tant que invalide)
func AskChoiceIndex(max int) int {
	for {
		s := AskChoice()
		n, err := strconv.Atoi(s)
		if err == nil && n >= 1 && n <= max {
			return n
		}
		fmt.Println("❌ Choix invalide.")
	}
}
