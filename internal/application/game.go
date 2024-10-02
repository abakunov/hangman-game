package application

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"hangman-game/config"
	"hangman-game/internal/domain"
	"math/big"
	"os"
	"strings"
)

// Scanner for obtaining user input.
var Scanner = bufio.NewScanner(os.Stdin)

// RunGame starts the game process.
func RunGame(categories map[string][]string) {
	PrintWelcomeMessage()
	PrintRules()

	category := getCategory(categories)
	difficulty := getDifficulty()
	attemptsLeft := getAttemptsByDifficulty(difficulty)

	word := getRandomWord(category, categories)
	game := domain.NewHangmanGame(word, attemptsLeft)

	for !game.IsOver() {
		stageIndex := len(config.HangmanStages) - game.AttemptsLeft - 1
		if stageIndex < 0 {
			stageIndex = 0
		}

		PrintGameStatus(config.HangmanStages[stageIndex], game.GetGuessedWord(), game.AttemptsLeft, game.WrongGuesses)

		letter := PromptLetter()
		if len(letter) != 1 {
			PrintMessage("Please enter a single letter.")
			continue
		}

		if game.Guess(rune(letter[0])) {
			PrintMessage("Correct letter!")
		} else {
			PrintMessage("Incorrect letter.")
		}
	}

	if game.HasWon() {
		PrintMessage("Congratulations! You guessed the word: " + game.Word)
	} else {
		PrintMessage(config.HangmanStages[len(config.HangmanStages)-1])
		PrintMessage("You lost. The word was: " + game.Word)
	}
}

func getCategory(categories map[string][]string) string {
	category := PromptCategory(categories)
	if _, exists := categories[category]; !exists || category == "" {
		PrintMessage("A random category has been selected.")

		category = getRandomKey(categories)
	}

	return category
}

func getDifficulty() string {
	return PromptDifficulty()
}

func getAttemptsByDifficulty(difficulty string) int {
	switch difficulty {
	case "easy":
		return 10
	case "medium":
		return 7
	case "hard":
		return 5
	default:
		PrintMessage("A random difficulty level has been selected.")
		return 7
	}
}

// getRandomKey returns a random key from the map.
func getRandomKey(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys[randomIndex(len(keys))]
}

// randomIndex returns a random index for selecting an element.
func randomIndex(maxIndex int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(maxIndex)))
	if err != nil {
		return 0
	}

	return int(n.Int64())
}

// getRandomWord returns a random word from the selected category.
func getRandomWord(category string, categories map[string][]string) string {
	words := categories[category]
	return words[randomIndex(len(words))]
}

// PromptCategory prompts the user to choose a category.
func PromptCategory(categories map[string][]string) string {
	fmt.Println("Choose a category:")

	for category := range categories {
		fmt.Println(category)
	}

	fmt.Print("Enter a category: ")

	Scanner.Scan()

	return strings.ToLower(Scanner.Text())
}

// PromptDifficulty prompts the user to choose a difficulty level.
func PromptDifficulty() string {
	fmt.Println("Choose a difficulty level (easy, medium, hard) or press Enter for random selection:")
	Scanner.Scan()

	return strings.ToLower(Scanner.Text())
}

// PromptLetter prompts the user to enter a letter.
func PromptLetter() string {
	fmt.Print("Enter a letter: ")

	if !Scanner.Scan() {
		PrintMessage("\nInput canceled. The game is over.")
		os.Exit(0)
	}

	return strings.ToLower(Scanner.Text())
}

// PrintMessage outputs a message to the screen.
func PrintMessage(message string) {
	fmt.Println(message)
}

// PrintGameStatus outputs the current game status.
func PrintGameStatus(stage, guessedWord string, attemptsLeft int, wrongGuesses []rune) {
	fmt.Println("\n" + stage)
	fmt.Println("Guessed word:", guessedWord)
	fmt.Printf("Attempts left: %d\n", attemptsLeft)
	fmt.Printf("Wrong letters: %c\n", wrongGuesses)
}
