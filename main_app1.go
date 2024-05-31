package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// App 1 - Playing with the App
func main() {
	pickSelection := "(1) rock - (2) paper - (3) scissors"
	botUserAnswers := [3]string{"rock", "paper", "scissors"}
	isGameRunning := true
	fmt.Println("GAME: rock paper scissors")
	fmt.Println("Hint: Input the keyword 'exit' in order to close the game")

	for isGameRunning {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Select your pick by typing in the word: ( " + pickSelection + " ): ")
		answer, _ := reader.ReadString('\n')
		text := strings.TrimSpace(answer)
		if text == "exit" {
			isGameRunning = false
			fmt.Println("Thanks for playing rock paper scissors!")
			fmt.Println("Exit game..")
			continue
		} else if text == "" {
			fmt.Println("Empty input detected. Please type in your pick: " + pickSelection)
		} else if text != "rock" || text != "paper" || text != "scissors" {
			fmt.Println("Please specify a valid pick: " + pickSelection)
		} else {
			randomPickNumber := rand.Intn(len(botUserAnswers))
			botUserPick := botUserAnswers[randomPickNumber]

			fmt.Println("Your pick was: " + text)
			fmt.Println("The bot user's pick was: " + botUserPick)
			looseText := "You loose."
			winText := "You win."

			if botUserPick == "paper" && text == "rock" {
				fmt.Println(looseText)
			} else if botUserPick == "paper" && text == "scissors" {
				fmt.Println(winText)
			} else if botUserPick == "rock" && text == "paper" {
				fmt.Println(winText)
			} else if botUserPick == "rock" && text == "scissors" {
				fmt.Println(looseText)
			} else if botUserPick == "scissors" && text == "paper" {
				fmt.Println(looseText)
			} else if botUserPick == "scissors" && text == "rock" {
				fmt.Println(winText)
			} else {
				fmt.Println("A Draw!")
			}

		}

	}

}
