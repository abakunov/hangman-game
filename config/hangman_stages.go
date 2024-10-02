package config

// HangmanStages contains the ASCII art for the hangman stages.
var HangmanStages = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
      |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`}
