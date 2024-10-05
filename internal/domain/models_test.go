package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hangman-game/internal/domain"
)

func TestNewHangmanGame(t *testing.T) {
	tests := []struct {
		name                 string
		word                 string
		attemptsLeft         int
		expectedWord         string
		expectedGuessedWord  []rune
		expectedAttemptsLeft int
	}{
		{
			name:                 "Base case with regular word",
			word:                 "hangman",
			attemptsLeft:         5,
			expectedWord:         "hangman",
			expectedGuessedWord:  []rune{'_', '_', '_', '_', '_', '_', '_'},
			expectedAttemptsLeft: 5,
		},
		{
			name:                 "Empty word",
			word:                 "",
			attemptsLeft:         5,
			expectedWord:         "",
			expectedGuessedWord:  []rune{},
			expectedAttemptsLeft: 5,
		},
		{
			name:                 "Zero attempts left",
			word:                 "hangman",
			attemptsLeft:         0,
			expectedWord:         "hangman",
			expectedGuessedWord:  []rune{'_', '_', '_', '_', '_', '_', '_'},
			expectedAttemptsLeft: 0,
		},
		{
			name:                 "Negative attempts left",
			word:                 "hangman",
			attemptsLeft:         -5,
			expectedWord:         "hangman",
			expectedGuessedWord:  []rune{'_', '_', '_', '_', '_', '_', '_'},
			expectedAttemptsLeft: -5,
		},
	}

	// Проходим по каждому тесту
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызов функции
			game := domain.NewHangmanGame(tt.word, tt.attemptsLeft)

			// Проверки
			assert.Equal(t, tt.expectedWord, game.Word, "Word must be initialized")
			assert.Equal(t, tt.expectedGuessedWord, game.GuessedWord, "GuessedWord must contain symbols '-'")
			assert.Equal(t, tt.expectedAttemptsLeft, game.AttemptsLeft, "Amount of attempts must be correct")
			assert.Empty(t, game.WrongGuesses, "WrongGuesses must be empty on a start")
		})
	}
}

func TestHangmanGame_Guess(t *testing.T) {
	tests := []struct {
		name   string
		letter rune
	}{
		{
			name:   "Correct guess",
			letter: 'h',
		},
		{
			name:   "Incorrect guess",
			letter: 'z',
		},
		{
			name:   "Incorrect format",
			letter: '1',
		},
	}

	g := domain.NewHangmanGame("hangman", 5)
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Вызов функции
			correct := g.Guess(testCase.letter)

			if testCase.name == "Correct guess" {
				assert.True(t, correct, "Guess must be correct")
			} else {
				assert.False(t, correct, "Guess must be incorrect")
			}
		})
	}
}

func TestHangmanGame_IsOver(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		attempts int
		guessed  []rune
		expected bool
	}{
		{
			name:     "Base case",
			word:     "hangman",
			attempts: 5,
			guessed:  []rune{'h', 'a', 'n', 'g', 'm'},
			expected: false,
		},
		{
			name:     "No attempts left",
			word:     "hangman",
			attempts: 0,
			guessed:  []rune{'h', 'a', 'n', 'g', 'm'},
			expected: true,
		},
		{
			name:     "Word is guessed",
			word:     "hangman",
			attempts: 5,
			guessed:  []rune{'h', 'a', 'n', 'g', 'm', 'a', 'n'},
			expected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			game := domain.NewHangmanGame(testCase.word, testCase.attempts)
			game.GuessedWord = testCase.guessed
			assert.Equal(t, testCase.expected, game.IsOver(), "Game must be over")
		})
	}
}

func TestHangmanGame_HasWon(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		guessed  []rune
		expected bool
	}{
		{
			name:     "Base case",
			word:     "hangman",
			guessed:  []rune{'h', 'a', 'n', 'g', 'm', 'a', 'n'},
			expected: true,
		},
		{
			name:     "Not guessed",
			word:     "hangman",
			guessed:  []rune{'h', 'a', 'n', 'g', 'm'},
			expected: false,
		},
		{
			name:     "Empty word",
			word:     "",
			guessed:  []rune{},
			expected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			game := domain.NewHangmanGame(testCase.word, 5)
			game.GuessedWord = testCase.guessed
			assert.Equal(t, testCase.expected, game.HasWon(), "Game must be won")
		})
	}
}

func TestHangmanGame_GetGuessedWord(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		guessed  []rune
		expected string
	}{
		{
			name:     "Base case",
			word:     "hangman",
			guessed:  []rune{'h', 'a', 'n', 'g', 'm', 'a', 'n'},
			expected: "hangman",
		},
		{
			name:     "Empty word",
			word:     "",
			guessed:  []rune{},
			expected: "",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			game := domain.NewHangmanGame(testCase.word, 5)
			game.GuessedWord = testCase.guessed
			assert.Equal(t, testCase.expected, game.GetGuessedWord(), "Guessed word must be correct")
		})
	}
}
