package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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

func revealTwoLetters(word string) map[string]bool {
	guessedChars := make(map[string]bool)
	wordChars := []rune(word)
	uniqueChars := make(map[string]bool)

	for _, char := range wordChars {
		uniqueChars[string(char)] = true
	}

	var uniqueCharSlice []string
	for char := range uniqueChars {
		uniqueCharSlice = append(uniqueCharSlice, char)
	}

	numToReveal := 2
	if len(uniqueCharSlice) < 2 {
		numToReveal = len(uniqueCharSlice)
	}

	rand.Shuffle(len(uniqueCharSlice), func(i, j int) {
		uniqueCharSlice[i], uniqueCharSlice[j] = uniqueCharSlice[j], uniqueCharSlice[i]
	})

	for i := 0; i < numToReveal; i++ {
		guessedChars[uniqueCharSlice[i]] = true
	}

	return guessedChars
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
	execDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	wordsPath := filepath.Join(execDir, "words.txt")

	secretWord, err := WordScanner(wordsPath)
	if err != nil {
		fmt.Println("Error loading words:", err)
		return
	}

	secretWord = strings.ToLower(secretWord)

	guessedChars := revealTwoLetters(secretWord)
	mistakeCount := 0

	fmt.Println("Welcome to Hangman! Two letters have been revealed to help you start.")
	fmt.Println("Press '0' anytime to quit the game.")

	for mistakeCount < maxMistakes {
		ShowGallows(mistakeCount)
		showProgress(secretWord, guessedChars)
		fmt.Println("Guessed characters:", listGuessedChars(guessedChars))
		fmt.Printf("Mistakes: %d / %d\n", mistakeCount, maxMistakes)

		fmt.Print("Guess a character or the whole word: ")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToLower(guess))

		if guess == "0" {
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
		} else {
			if guess == secretWord {
				fmt.Println("\nCongratulations! You've guessed the word:", secretWord)
				ShowGallows(mistakeCount)
				return
			} else if guess != "" {
				fmt.Println("Wrong word guess!")
				mistakeCount++
			} else {
				fmt.Println("Invalid input.")
			}
		}

		if isWordDiscovered(secretWord, guessedChars) {
			fmt.Println("\nWell done! You've uncovered the word:", secretWord)
			ShowGallows(mistakeCount)
			return
		}
	}

	fmt.Println("\nGame over. The word was:", secretWord)
	ShowGallows(mistakeCount)
}
