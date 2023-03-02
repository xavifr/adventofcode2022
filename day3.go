package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prepareMap() map[rune]int {
	ret := map[rune]int{}
	cnt := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		ret[ch] = cnt
		cnt++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		ret[ch] = cnt
		cnt++
	}

	return ret
}

func main() {
	day3_part1()
	day3_part2()
}

func day3_part1() {
	scoreMap := prepareMap()

	fmt.Println("DAY3: PART 1")
	file, err := os.Open("day3.input")
	if err != nil {
		fmt.Printf("Error downloading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText := scanner.Text()
		if len(inputText)%2 != 0 {
			fmt.Printf("INPUT IS NOT EVEN!!")
			return
		}

		firstCompartment := inputText[:len(inputText)/2]
		secondCompartment := inputText[len(inputText)/2:]

		for _, fl := range firstCompartment {
			if strings.Contains(secondCompartment, string(fl)) {
				score += scoreMap[fl]
				break
			}
		}
	}

	fmt.Printf("SCORE IS: %d\n", score)
}

func day3_part2() {
	scoreMap := prepareMap()

	fmt.Println("DAY3: PART 2")
	file, err := os.Open("day3.input")
	if err != nil {
		fmt.Printf("Error downloading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elf1 := scanner.Text()
		if !scanner.Scan() {
			break
		}
		elf2 := scanner.Text()
		if !scanner.Scan() {
			break
		}
		elf3 := scanner.Text()

		for _, fl := range elf1 {
			if strings.Contains(elf2, string(fl)) && strings.Contains(elf3, string(fl)) {
				score += scoreMap[fl]
				break
			}
		}
	}

	fmt.Printf("SCORE IS: %d\n", score)
}
