package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day4_part1()
	day4_part2()
}

func day4_part1() {
	fmt.Println("DAY4: PART 1")
	file, err := os.Open("day4.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	re, _ := regexp.Compile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		m := re.FindStringSubmatch(scanner.Text())
		if m == nil || len(m) == 0 {
			fmt.Printf("Error in put madafaka: %s\n", scanner.Text())
			return
		}

		mi := make([]int, 5)
		mi[1], _ = strconv.Atoi(m[1])
		mi[2], _ = strconv.Atoi(m[2])
		mi[3], _ = strconv.Atoi(m[3])
		mi[4], _ = strconv.Atoi(m[4])

		if (mi[1] >= mi[3] && mi[2] <= mi[4]) || (mi[3] >= mi[1] && mi[4] <= mi[2]) {
			//	fmt.Printf("RANGE %d-%d IS CONTAINED OR CONTAINS %d-%d\n", mi[1], mi[2], mi[3], mi[4])
			count++
		}
	}

	fmt.Println(count)
}
func day4_part2() {
	fmt.Println("DAY4: PART 2")
	file, err := os.Open("day4.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	re, _ := regexp.Compile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())
		if m == nil || len(m) == 0 {
			fmt.Printf("Error in put madafaka: %s\n", scanner.Text())
			return
		}

		mi := make([]int, 5)
		mi[1], _ = strconv.Atoi(m[1])
		mi[2], _ = strconv.Atoi(m[2])
		mi[3], _ = strconv.Atoi(m[3])
		mi[4], _ = strconv.Atoi(m[4])

		if (mi[3] <= mi[2] && mi[4] >= mi[1]) || (mi[1] <= mi[4] && mi[2] > mi[3]) || (mi[1] >= mi[3] && mi[2] <= mi[4]) || (mi[3] >= mi[1] && mi[4] <= mi[2]) {
			fmt.Printf("RANGE %d-%d OVERLAPS OR IS OVERLAPED WITH %d-%d\n", mi[1], mi[2], mi[3], mi[4])
			count++
		}
	}

	fmt.Println(count)
}
