package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day1.input")
	if err != nil {
		fmt.Printf("Error downloading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	caloriesPerElf := map[int]int{}

	maxCaloriesCarried := 0
	currentCaloriesCarried := 0
	currentElf := 1
	for scanner.Scan() {
		buff := scanner.Text()

		if buff == "" {
			if currentCaloriesCarried > maxCaloriesCarried {
				maxCaloriesCarried = currentCaloriesCarried
			}

			caloriesPerElf[currentElf] = currentCaloriesCarried
			currentElf++
			currentCaloriesCarried = 0
			continue
		}

		i, e := strconv.Atoi(scanner.Text())
		if e != nil {
			fmt.Printf("Error at calories input: %s\n", scanner.Text())
			continue
		}
		currentCaloriesCarried += i
	}

	if currentCaloriesCarried > 0 {
		if currentCaloriesCarried > maxCaloriesCarried {
			maxCaloriesCarried = currentCaloriesCarried
		}

		caloriesPerElf[currentElf] = currentCaloriesCarried
		currentCaloriesCarried = 0
	}

	if currentElf < 3 {
		fmt.Printf("Input length is less than 3 :(")
		return
	}

	fmt.Printf("The number of elves is: %d\n", currentElf)

	keys := make([]int, 0, len(caloriesPerElf))
	for key := range caloriesPerElf {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return caloriesPerElf[keys[i]] > caloriesPerElf[keys[j]]
	})

	if len(keys) < 3 {
		fmt.Printf("Input length is less than 3 :(")
		return
	}

	fmt.Printf("The elf with most calories carried is number %d with %d calores\n", keys[0], caloriesPerElf[keys[0]])
	fmt.Printf("The elf with second most calories carried is number %d with %d calores\n", keys[1], caloriesPerElf[keys[1]])
	fmt.Printf("The elf with third most calories carried is number %d with %d calores\n", keys[2], caloriesPerElf[keys[2]])

	fmt.Printf("The number of calories carried by those elves is: %d\n", caloriesPerElf[keys[0]]+caloriesPerElf[keys[1]]+caloriesPerElf[keys[2]])
}
