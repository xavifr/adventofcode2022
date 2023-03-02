package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day10_part1()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	day10_part2()
}

func day10_part1() {
	fmt.Println("DAY10: PART 1")

	file, err := os.Open("day10.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	reCmdAdd, _ := regexp.Compile(`^addx (-?\d+)$`)
	reCmdNoop, _ := regexp.Compile(`^noop$`)

	cicle := 0
	registerX := 1
	signal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if reCmdAdd.MatchString(scanner.Text()) {
			match := reCmdAdd.FindStringSubmatch(scanner.Text())
			valueAdded, _ := strconv.Atoi(match[1])
			//fmt.Printf("@%03d ADD: %d + %d\n", cicle, registerX, valueAdded)

			intBefore := (cicle + 20) / 40
			intAfter := (cicle + 22) / 40
			//fmt.Printf("  %d <=> %d\n", intBefore, intAfter)
			if intBefore != intAfter {
				//fmt.Printf("    RECORD SIGNAL\n")
				if (cicle+22)%40 == 0 {
					//fmt.Printf("        With complete cicle (%d, %d)\n", cicle+2, registerX)
					signal += (cicle + 2) * registerX
				} else {
					//fmt.Printf("        With semi cycle (%d, %d)\n", cicle+((cicle+22)%40), registerX)
					signal += (cicle + ((cicle + 22) % 40)) * registerX
				}
				//fmt.Printf("SIGNAL IS %d\n", signal)
			}
			registerX += valueAdded
			cicle += 2
		} else if reCmdNoop.MatchString(scanner.Text()) {
			//fmt.Printf("@%03d NOOP: %d\n", cicle, registerX)

			cicle += 1
			if (cicle+20)%40 == 0 {
				//fmt.Printf("    RECORD SIGNAL\n")
				//fmt.Printf("        With complete cicle (%d, %d)\n", cicle+2, registerX)
				signal += cicle * registerX
				//fmt.Printf("SIGNAL IS %d\n", signal)
			}
		}

	}

	fmt.Printf("Signal is %d and registerX %d\n", signal, registerX)
}
func day10_part2() {
	fmt.Println("DAY10: PART 2")

	file, err := os.Open("day10.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	reCmdAdd, _ := regexp.Compile(`^addx (-?\d+)$`)
	reCmdNoop, _ := regexp.Compile(`^noop$`)

	cycle := 0
	registerX := 1
	sprite := 1
	var output []rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if reCmdAdd.MatchString(scanner.Text()) {
			match := reCmdAdd.FindStringSubmatch(scanner.Text())
			valueAdded, _ := strconv.Atoi(match[1])

			if math.Abs(float64(sprite-(cycle%40))) <= 1 {
				output = append(output, '#')
			} else {
				output = append(output, '.')
			}

			cycle += 1
			if math.Abs(float64(sprite-(cycle%40))) <= 1 {
				output = append(output, '#')
			} else {
				output = append(output, '.')
			}
			cycle += 1
			registerX += valueAdded
			sprite = registerX
		} else if reCmdNoop.MatchString(scanner.Text()) {
			fmt.Printf("@%03d NOOP: %d\n", cycle, registerX)

			if math.Abs(float64(sprite-(cycle%40))) <= 1 {
				output = append(output, '#')
			} else {
				output = append(output, '.')
			}
			cycle += 1
		}

	}

	fmt.Printf("%+v\n", output)
	for i, c := range output {
		if i%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", c)
	}
}
