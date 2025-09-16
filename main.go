package main

import (
	"math/rand"
	"time"

	"main/game"
	"main/intro"
)

func main() {
	// seed random
	rand.Seed(time.Now().UnixNano())

	intro.ShowIntro()
	game.StartGame()
}
