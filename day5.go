package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(el rune) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return ' ', false
	}

	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return el, true
}

func (s *Stack) MPop(qty int) ([]rune, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	els := (*s)[len(*s)-qty:]
	*s = (*s)[:len(*s)-qty]

	return els, true
}

func (s *Stack) MPush(els []rune) {
	for _, el := range els {
		*s = append(*s, el)
	}
}

func (s *Stack) Head() (rune, bool) {
	if s.IsEmpty() {
		return ' ', false
	}

	el := (*s)[len(*s)-1]

	return el, true
}

func main() {
	day5_part1()
	day5_part2()
}

func day5_part1() {
	fmt.Println("DAY5: PART 1")
	file, err := os.Open("day5.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}
	//WSDZFHLNN

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := day5_initialize_stack(scanner)
	stacksKeys := make([]int, 0)
	for k, _ := range stacks {
		stacksKeys = append(stacksKeys, k)
	}

	sort.Ints(stacksKeys)

	fmt.Printf("%+v\n", stacks)
	for _, k := range stacksKeys {
		c, _ := stacks[k].Head()
		fmt.Printf("%d: %c\n", k, c)
	}
	fmt.Println("")

	reMove, _ := regexp.Compile(`^move (\d+) from (\d+) to (\d+)$`)
	for scanner.Scan() {
		if !reMove.MatchString(scanner.Text()) {
			continue
		}

		match := reMove.FindStringSubmatch(scanner.Text())
		qty, _ := strconv.Atoi(match[1])
		src, _ := strconv.Atoi(match[2])
		dst, _ := strconv.Atoi(match[3])

		//fmt.Printf("MOVE %d FROM %d to %d\n", qty, src, dst)
		srcStack := stacks[src]
		dstStack := stacks[dst]

		for i := 0; i < qty; i++ {
			crate, e := srcStack.Pop()
			if !e {
				fmt.Printf("EEEEEEEEE EMPTYU MADAFAKA %+v\n", srcStack)
				return
			}
			dstStack.Push(crate)
		}
	}

	for _, k := range stacksKeys {
		c, _ := stacks[k].Head()
		fmt.Printf("%c", c)
	}

	fmt.Println("")
}
func day5_part2() {
	fmt.Println("DAY5: PART 2")
	file, err := os.Open("day5.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := day5_initialize_stack(scanner)
	stacksKeys := make([]int, 0)
	for k, _ := range stacks {
		stacksKeys = append(stacksKeys, k)
	}

	sort.Ints(stacksKeys)

	fmt.Printf("%+v\n", stacks)
	for _, k := range stacksKeys {
		c, _ := stacks[k].Head()
		fmt.Printf("%d: %c\n", k, c)
	}
	fmt.Println("")

	reMove, _ := regexp.Compile(`^move (\d+) from (\d+) to (\d+)$`)
	for scanner.Scan() {
		if !reMove.MatchString(scanner.Text()) {
			continue
		}

		match := reMove.FindStringSubmatch(scanner.Text())
		qty, _ := strconv.Atoi(match[1])
		src, _ := strconv.Atoi(match[2])
		dst, _ := strconv.Atoi(match[3])

		//fmt.Printf("MOVE %d FROM %d to %d\n", qty, src, dst)
		srcStack := stacks[src]
		dstStack := stacks[dst]

		crates, e := srcStack.MPop(qty)
		if !e {
			fmt.Printf("EEEEEEEEE EMPTYU MADAFAKA %+v\n", srcStack)
			return
		}
		dstStack.MPush(crates)
	}

	for _, k := range stacksKeys {
		c, _ := stacks[k].Head()
		fmt.Printf("%c", c)
	}

	fmt.Println("")
}

func day5_initialize_stack(scanner *bufio.Scanner) map[int]*Stack {
	stacks := map[int]*Stack{}

	var stackStatus []string

	ReDetect, _ := regexp.Compile(`^\s+(\d+)`)
	ReSplit, _ := regexp.Compile(`\s+`)

	for scanner.Scan() {
		if !ReDetect.MatchString(scanner.Text()) {
			stackStatus = append(stackStatus, scanner.Text())
			continue
		}

		stacksIndices := ReSplit.Split(scanner.Text(), -1)
		for _, i := range stacksIndices {
			iS, e := strconv.Atoi(i)
			if e != nil {
				continue
			}
			stacks[iS] = new(Stack)
		}

		break
	}

	for i, j := 0, len(stackStatus)-1; i < j; i, j = i+1, j-1 {
		stackStatus[i], stackStatus[j] = stackStatus[j], stackStatus[i]
	}

	for _, line := range stackStatus {
		fmt.Printf("STATUS: %s\n", line)
		for i, _ := range stacks {
			offset := i + (i-1)*3
			fmt.Printf("STACK: %d, OFFSET: %d, #LINE: %d\n", i, offset, len(line))
			if offset >= len(line) {
				continue

			}
			crate := line[offset]

			if crate != ' ' {
				fmt.Printf("  APPEND CRATE: %c\n", rune(crate))
				stacks[i].Push(rune(crate))
			}
		}
	}

	return stacks
}
