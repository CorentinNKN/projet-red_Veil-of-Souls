package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// AskChoice lit une entrée clavier (trim + toLower) pour uniformiser
func AskChoice() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// SleepSeconds fait une pause de X secondes
func SleepSeconds(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// Pause attend que l'utilisateur appuie sur Entrée (utile après un combat ou un écran de victoire)
func Pause() {
	fmt.Println("\nAppuyez sur Entrée pour continuer...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
