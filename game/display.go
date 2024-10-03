package game

import "fmt"

func ShowMainMenu() {

	fmt.Println("\n                                        === Main Menu ===")
	fmt.Println("\n                                         1. Start The Game")
	fmt.Println("\n                                          2. Rules")
	fmt.Println("\n                                           3. Quit")

	var choice int

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		StartGame()
	case 2:
		ShowRules()
	case 3:
		fmt.Println("See You Next Time.....")
		return
	default:
		fmt.Println("Invalid Choice !!")
	}

}

func ShowRules() {
	fmt.Println("\n                === Rules Menu ===")
	fmt.Println("The rules follow the traditional Hangman game format.  ")
	fmt.Println("You have a word to guess with a total of 8 attempts.")
	fmt.Println("Additionally, you have the option to guess the entire word at once,")
	fmt.Println("but if your guess is incorrect, you will lose 2 attempts.  ")

	ShowMainMenu()
}

func ShowGallows(mistakes int) {
	gallowsStages := []string{
		`
        +---+
        |   |
            |
            |
            |
            |
        =========`,
		`
        +---+
        |   |
        O   |
            |
            |
            |
        =========`,
		`
        +---+
        |   |
        O   |
        |   |
            |
            |
        =========`,
		`
        +---+
        |   |
        O   |
       /|   |
            |
            |
        =========`,
		`
        +---+
        |   |
        O   |
       /|\  |
            |
            |
        =========`,
		`
        +---+
        |   |
        O   |
       /|\  |
       /    |
            |
        =========`,
		`
        +---+
        |   |
        O   |
       /|\  |
       / \  |
            |
        =========`,
		`
        +---+
        |   |
        O   |
       /|\  |
       / \  |
       |    |
        =========`,
	}

	if mistakes < len(gallowsStages) {
		fmt.Println(gallowsStages[mistakes])
	}
}
