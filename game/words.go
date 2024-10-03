package game

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func WordScanner(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))], nil
}
