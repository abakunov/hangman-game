package domain

// HangmanGame represents the state of the game.
type HangmanGame struct {
	Word         string // The word to guess
	GuessedWord  []rune // The current state of the guessed word
	AttemptsLeft int    // The number of attempts left
	WrongGuesses []rune // The list of wrong guesses
}

// NewHangmanGame creates a new game with the given word and attempts left.
func NewHangmanGame(word string, attemptsLeft int) *HangmanGame {
	guessedWord := make([]rune, len(word))
	for i := range guessedWord {
		guessedWord[i] = '_'
	}

	return &HangmanGame{
		Word:         word,
		GuessedWord:  guessedWord,
		AttemptsLeft: attemptsLeft,
		WrongGuesses: []rune{},
	}
}

// Guess checks the entered letter and updates the game state.
func (g *HangmanGame) Guess(letter rune) bool {
	correct := false

	for i, l := range g.Word {
		if l == letter {
			g.GuessedWord[i] = letter
			correct = true
		}
	}

	if !correct {
		g.AttemptsLeft--
		g.WrongGuesses = append(g.WrongGuesses, letter)
	}

	return correct
}

// IsOver checks if the game is over.
func (g *HangmanGame) IsOver() bool {
	return g.AttemptsLeft == 0 || g.HasWon()
}

// HasWon checks if the player has won.
func (g *HangmanGame) HasWon() bool {
	return string(g.GuessedWord) == g.Word
}

// GetGuessedWord returns the current state of the guessed word.
func (g *HangmanGame) GetGuessedWord() string {
	return string(g.GuessedWord)
}
