package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskChoice() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Votre choix : ")
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}
