package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxMistakes = 8

func showProgress(word string, guessedChars map[string]bool) {
	for _, char := range word {
		if guessedChars[string(char)] {
			fmt.Print(string(char), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func isWordDiscovered(word string, guessedChars map[string]bool) bool {
	for _, char := range word {
		if !guessedChars[string(char)] {
			return false
		}
	}
	return true
}

func listGuessedChars(guessedChars map[string]bool) string {
	var chars []string
	for char := range guessedChars {
		chars = append(chars, char)
	}
	return strings.Join(chars, ", ")
}

func StartGame() {
	secretWord, err := WordScanner("Documents/Ynov 2425/Module Immersion/Hangman-gowords.txt") // Here you put the file path of the text file
	if err != nil {
		fmt.Println("Error loading words:", err)
		return
	}

	guessedChars := make(map[string]bool)
	mistakeCount := 0

	fmt.Println("Press 'Q' anytime to quit the game.")

	for mistakeCount < maxMistakes {
		showGallows(mistakeCount)
		showProgress(secretWord, guessedChars)
		fmt.Println("Guessed characters:", listGuessedChars(guessedChars))
		fmt.Printf("Mistakes: %d / %d\n", mistakeCount, maxMistakes)

		fmt.Print("Guess a character or the whole word: ")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToLower(guess))

		if guess == "q" {
			fmt.Println("Game aborted.")
			return
		}

		if len(guess) == 1 {
			if _, exists := guessedChars[guess]; exists {
				fmt.Println("You've already guessed this character.")
				continue
			}

			guessedChars[guess] = true
			if strings.Contains(secretWord, guess) {
				fmt.Println("Good guess!")
			} else {
				fmt.Println("Wrong guess.")
				mistakeCount++
			}
		} else if len(guess) == len(secretWord) {
			if guess == secretWord {
				fmt.Println("Congratulations! You've guessed the word:", secretWord)
				return
			} else {
			}
		} else {
			fmt.Println("Invalid input.")
		}

		if isWordDiscovered(secretWord, guessedChars) {
			fmt.Println("Well done! You've uncovered the word:", secretWord)
			return
		}
	}

	fmt.Println("Game over. The word was:", secretWord)
}
