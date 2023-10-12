package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	words := strings.Split(alphabet, "")
	randomWord := getRandomWord(words)
	guesses := make([]string, 0)
	maxAttempts := 6
	attempts := 0

	fmt.Println("Bienvenue au jeu du pendu!")
	displayWord(randomWord, guesses)

	for {
		var guess string
		fmt.Print("Devinez une lettre: ")
		fmt.Scanln(&guess)

		if strings.Contains(randomWord, guess) {
			fmt.Println("Bonne devinette!")
			guesses = append(guesses, guess)
		} else {
			fmt.Println("Mauvaise devinette!")
			attempts++
		}

		displayWord(randomWord, guesses)

		if isWinner(randomWord, guesses) {
			fmt.Println("Félicitations, vous avez gagné!")
			break
		}

		if attempts >= maxAttempts {
			fmt.Println("Désolé, vous avez perdu. Le mot était", randomWord)
			break
		}
	}
}

func getRandomWord(words []string) string {
	rand.Seed(42) // Pour rendre la génération de nombres aléatoires reproductible
	return words[rand.Intn(len(words))]
}

func displayWord(word string, guesses []string) {
	displayedWord := ""
	for _, letter := range word {
		if contains(guesses, string(letter)) {
			displayedWord += string(letter) + " "
		} else {
			displayedWord += "_ "
		}
	}
	fmt.Println(displayedWord)
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

func isWinner(word string, guesses []string) bool {
	for _, letter := range word {
		if !contains(guesses, string(letter)) {
			return false
		}
	}
	return true
}
