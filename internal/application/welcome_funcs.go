package application

func PrintWelcomeMessage() {
	PrintMessage("========================================")
	PrintMessage("     Welcome to the 'Hangman' game!")
	PrintMessage("========================================")
}

func PrintRules() {
	PrintMessage("Game Rules:")
	PrintMessage("1. Choose a word category and difficulty level.")
	PrintMessage("2. You have a limited number of attempts to guess letters.")
	PrintMessage("3. Guess letters one at a time. If a letter is in the word, it will be revealed.")
	PrintMessage("4. If the letter is not present, the number of attempts decreases.")
	PrintMessage("5. The game ends when you guess the word or run out of attempts.")
	PrintMessage("Good luck!\n")
}
