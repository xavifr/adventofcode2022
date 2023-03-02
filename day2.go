package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type figures struct {
	Name   string
	Score  int
	WinsTo string
}

type figures2 struct {
	Name       string
	Score      int
	WinsTo     string
	LoosesWith string
}

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("day2.input")
	if err != nil {
		fmt.Printf("Error downloading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	inputs := map[string]figures{"X": {"Rock", 1, "Scissors"}, "Y": {Name: "Paper", Score: 2, WinsTo: "Rock"}, "Z": {Name: "Scissors", Score: 3, WinsTo: "Paper"}}
	reactions := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray := strings.Split(scanner.Text(), " ")
		if len(inputArray) != 2 {
			fmt.Printf("INPUT NOT VALID: %s\n", scanner.Text())
			return
		}

		input := inputs[inputArray[1]]
		reaction := reactions[inputArray[0]]

		//fmt.Printf("FIGHT! %s <=> %s: ", input.Name, reaction)
		score += input.Score
		if input.Name == reaction {
			score += 3
			//	fmt.Println("DRAW")
		} else if input.WinsTo == reaction {
			score += 6
			//	fmt.Println("DRAW")
		} else {
			//	fmt.Println("LOSS")
		}
	}

	fmt.Printf("PART1: FINAL SCORE IS: %d\n", score)

}
func part2() {
	file, err := os.Open("day2.input")
	if err != nil {
		fmt.Printf("Error downloading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	inputs := map[string]figures2{"A": {"Rock", 1, "C", "B"}, "B": {"Paper", 2, "A", "C"}, "C": {"Scissors", 3, "B", "A"}}

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray := strings.Split(scanner.Text(), " ")
		if len(inputArray) != 2 {
			fmt.Printf("INPUT NOT VALID: %s\n", scanner.Text())
			return
		}

		input := inputs[inputArray[0]]

		if inputArray[1] == "X" {
			reaction := inputs[input.WinsTo]
			score += reaction.Score + 0
			fmt.Printf("FIGHT! %s <=> %s: LOOSE\n", input.Name, reaction.Name)
		} else if inputArray[1] == "Y" {
			reaction := input
			score += reaction.Score + 3
			fmt.Printf("FIGHT! %s <=> %s: DRAW\n", input.Name, reaction.Name)
		} else if inputArray[1] == "Z" {
			reaction := inputs[input.LoosesWith]
			score += reaction.Score + 6
			fmt.Printf("FIGHT! %s <=> %s: WINS\n", input.Name, reaction.Name)
		} else {
			fmt.Println("UNKNOWN REACTION")
			continue
		}
	}

	fmt.Printf("PART2: FINAL SCORE IS: %d\n", score)

}
