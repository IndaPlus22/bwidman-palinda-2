// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"strconv"
)

const (
	star   = "The Fortune Teller"
	venue  = "Kiyamachi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("What do you want to know?")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go receive(questions, answers)
	go prediction(answers)
	go printAnswers(answers)

	return questions
}

func receive(questions <-chan string, answers chan<- string) {
	questionsAsked := 1
	for {
		go prophecy(<-questions, answers, questionsAsked)
		questionsAsked++
	}
}

func prediction(answers chan<- string) {
	for {
		time.Sleep(time.Duration(30+rand.Intn(20)) * time.Second)

		// Cook up some pointless nonsense.
		nonsense := []string{
			star + " will not only answer you but also make predictions, which one just got fulfilled.",
			"The day artificial intelligence gain authority, humans will regret their ingenuity.",
			"You will die in five minutes lmao",
		}
		randIdx := rand.Intn(len(nonsense))
		answers <- nonsense[randIdx]
	}
}

func printAnswers(answers <-chan string) {
	for {
		answer := <-answers
		// Print answer character by character with a short delay
		for _, char := range answer {
			fmt.Print(string(char))
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Println()
		fmt.Print(prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string, questionsAsked int) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	if strings.ToLower(question) == "what is the meaning of life?" {
		answer <- "Ah, life!..."
		return
	} else if strings.Contains(question, "rose-colored campus life") {
		answer <- "There is no such thing as a rose-colored campus life. There is nothing rose-colored in this world. Everything is all a bunch of colors mixed up, you see."
		return
	}

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}
	
	// Cook up some pointless nonsense.
	nonsense := []string{
		"You shall see it when the moon is at its darkest.",
		"You shall see it when the sun is at its highest.",
		"It's an opportunity you must grab and act on.",
		"",
	}
	randIdx := rand.Intn(len(nonsense))
	answer <- longestWord + 
		"... The answer is always dangling in front of your eyes. " + 
		nonsense[randIdx] +
		" OK that'll be " + strconv.Itoa(questionsAsked * 1000) + " yen."
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
