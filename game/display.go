package game

import "fmt"

func ShowMainMenu() {
	fmt.Println(`
                    ╔══════════════════════════════════════════════╗
                    ║                AWESOME HANGMAN               ║
                    ╚══════════════════════════════════════════════╝

                    ╭──────────────────────────────────────────────╮
                    │                                              │
                    │              [1]  START THE GAME             │
                    │                                              │
                    │              [2]  RULES                      │
                    │                                              │
                    │              [3]  QUIT                       │
                    │                                              │
                    ╰──────────────────────────────────────────────╯
`)

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		StartGame()
	case 2:
		ShowRules()
	case 3:
		fmt.Println(`
                    ╭──────────────────────────────────────────────╮
                    │            See You Next Time!                │
                    ╰──────────────────────────────────────────────╯
`)
		return
	default:
		fmt.Println(`
                    ╭──────────────────────────────────────────────╮
                    │            Invalid Choice!!                  │
                    ╰──────────────────────────────────────────────╯
`)
		ShowMainMenu()
	}
}

func ShowRules() {
	fmt.Println(`
                    ╔══════════════════════════════════════════════╗
                    ║                   RULES                      ║
                    ╚══════════════════════════════════════════════╝

                    ╭──────────────────────────────────────────────╮
                    │  • The rules follow the traditional          │
                    │    Hangman game format.                      │
                    │                                              │
                    │  • You have a word to guess with a           │
                    │    total of 8 attempts.                      │
                    │                                              │
                    │  • You can guess the entire word at once,    │
                    │    but if wrong, you lose 2 attempts.        │
                    ╰──────────────────────────────────────────────╯
`)
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
