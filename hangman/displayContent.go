package hangman

import "fmt"

func showMainMenu() {
	fmt.Println("\n                                        === Main Menu ===")
	fmt.Println("\n                                           1. Start")
	fmt.Println("\n                                           2. Rules")
	fmt.Println("\n                                           0. Quit")

	var choice int

	fmt.Scanln(choice)

	switch choice {
	case 1:
		startGame()
	case 2:
		showRules()
	case 0:
		fmt.Println("See You Next Time.....")

	default:
		fmt.Println("Invalid Choice !!")
		showMainMenu()

	}
}

func showRules() {
	fmt.Println("\n                   p-z                     === Rules Menu ===")
	fmt.Println("         The rules follow the traditional Hangman game format. You have a word to guess with a total of 7 attempts. Additionally, you have the option to guess the entire word at once, but if your guess is incorrect, you will lose 2 attempts.  ")
	fmt.Println("\n                                       0 To return to Main menu")
	var choice int
	fmt.Scanln(choice)

	if choice == 0 {
		showMainMenu()
	}
}
