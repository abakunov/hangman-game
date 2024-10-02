package main

import (
	"hangman-game/internal/application"
	"hangman-game/internal/infrastructure"
)

func main() {
	// Load words from the JSON file
	categories, err := infrastructure.LoadWordsFromFile("../../config/words.json")
	if err != nil {
		application.PrintMessage("Error loading words: " + err.Error())
		return
	}

	// Start the game
	application.RunGame(categories)
}
