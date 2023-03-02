package main

import (
	"bufio"
	"fmt"
	"github.com/Pramod-Devireddy/go-exprtk"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day11_part1()
	day11_part2()
}

type Monkey struct {
	Number              int
	Items               []int64
	OperationExpression string
	Expression          exprtk.GoExprtk
	OperationOperation  string
	OperationValue      string
	TestDivisibleBy     int
	TestPositive        int
	TestNegative        int
	ItemsInspected      int
}

func (m *Monkey) Reduce(divisible int64) {
	if len(m.Items) == 0 {
		return
	}

	for i, _ := range m.Items {
		if m.Items[i] > divisible {
			m.Items[i] %= divisible
		}
	}
}

func (m *Monkey) InspectItem(divide bool) (bool, int, int64) {
	if len(m.Items) == 0 {
		return false, 0, 0
	}
	item := m.Items[0]
	//fmt.Printf("    Monkey inspects item with wl %d\n", item)
	m.Items = m.Items[1:]
	m.ItemsInspected++

	m.Expression.SetDoubleVariableValue("old", float64(item))
	newItem := m.Expression.GetEvaluatedValue()
	//fmt.Printf("      WL is worried, now is %f\n", newItem)
	var newItemBored int64
	if divide {
		newItemBored = int64(math.Floor(newItem / 3))
	} else {
		newItemBored = int64(newItem)
	}
	/*
		newItem := item
		if m.OperationOperation == "+" {
			if m.OperationValue == "old" {
				newItem += item
			} else {
				optVal, _ := strconv.ParseInt(m.OperationValue, 10, 64)
				newItem += optVal
			}
		} else if m.OperationOperation == "*" {
			if m.OperationValue == "old" {
				newItem *= item
			} else {
				optVal, _ := strconv.ParseInt(m.OperationValue, 10, 64)
				newItem *= optVal
			}
		} else {
			fmt.Printf("Unkown operation %s\n", m.OperationOperation)
		}


		fmt.Printf("      WL is worried, now is %d\n", newItem)
		var newItemBored int64
		if divide {
			newItemBored = newItem / 3
		} else {
			newItemBored = newItem
		}

	*/
	//fmt.Printf("      Monkey gets bored, WL now is %d\n", newItemBored)

	if newItemBored%int64(m.TestDivisibleBy) == int64(0) {
		//fmt.Printf("      WL is divisible by %d, dispatched to %d\n", m.TestDivisibleBy, m.TestPositive)
		return true, m.TestPositive, newItemBored
	} else {
		//fmt.Printf("      WL is no divisible by %d, dispatched to %d\n", m.TestDivisibleBy, m.TestNegative)
		return true, m.TestNegative, newItemBored
	}
}

func day11_import_monkeys() map[int]Monkey {
	monkeys := map[int]Monkey{}
	file, err := os.Open("day11.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return nil
	}

	defer file.Close()

	reMonkeyStart, _ := regexp.Compile(`^Monkey (\d+):$`)
	reMonkeyItems, _ := regexp.Compile(`^\s+Starting items: (\d+(,\s+\d+)*)$`)
	reMonkeyOperation, _ := regexp.Compile(`^\s+Operation: new = (.*)$`)
	//reMonkeyOperation, _ := regexp.Compile(`^\s+Operation: new = old ([+*]) (old|\d+)$`)
	reMonkeyTest, _ := regexp.Compile(`^\s+Test: divisible by (\d+)$`)
	reMonkeyTestTrue, _ := regexp.Compile(`^\s+If true: throw to monkey (\d+)$`)
	reMonkeyTestFalse, _ := regexp.Compile(`^\s+If false: throw to monkey (\d+)$`)

	var currentMonkey Monkey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			if currentMonkey.TestDivisibleBy != 0 {
				monkeys[currentMonkey.Number] = currentMonkey
			}
		} else if reMonkeyStart.MatchString(scanner.Text()) {
			match := reMonkeyStart.FindStringSubmatch(scanner.Text())
			monkeyNumber, _ := strconv.Atoi(match[1])
			currentMonkey = Monkey{Number: monkeyNumber}
		} else if reMonkeyItems.MatchString(scanner.Text()) {
			match := reMonkeyItems.FindStringSubmatch(scanner.Text())
			items := strings.Split(match[1], ",")
			currentMonkey.Items = []int64{}
			for _, item := range items {
				itemNumber, _ := strconv.ParseInt(strings.TrimSpace(item), 10, 64)
				currentMonkey.Items = append(currentMonkey.Items, itemNumber)
			}
			/*} else if reMonkeyOperation.MatchString(scanner.Text()) {
			match := reMonkeyOperation.FindStringSubmatch(scanner.Text())
			currentMonkey.OperationOperation = match[1]
			currentMonkey.OperationValue = match[2]*/
		} else if reMonkeyOperation.MatchString(scanner.Text()) {
			match := reMonkeyOperation.FindStringSubmatch(scanner.Text())
			currentMonkey.OperationExpression = match[1]
			currentMonkey.Expression = exprtk.NewExprtk()
			currentMonkey.Expression.SetExpression(currentMonkey.OperationExpression)
			currentMonkey.Expression.AddDoubleVariable("old")
			currentMonkey.Expression.CompileExpression()

		} else if reMonkeyTest.MatchString(scanner.Text()) {
			match := reMonkeyTest.FindStringSubmatch(scanner.Text())
			divisibleBy, _ := strconv.Atoi(match[1])
			currentMonkey.TestDivisibleBy = divisibleBy
		} else if reMonkeyTestTrue.MatchString(scanner.Text()) {
			match := reMonkeyTestTrue.FindStringSubmatch(scanner.Text())
			targetMonkey, _ := strconv.Atoi(match[1])
			currentMonkey.TestPositive = targetMonkey
		} else if reMonkeyTestFalse.MatchString(scanner.Text()) {
			match := reMonkeyTestFalse.FindStringSubmatch(scanner.Text())
			targetMonkey, _ := strconv.Atoi(match[1])
			currentMonkey.TestNegative = targetMonkey
		}
	}

	if currentMonkey.TestDivisibleBy != 0 {
		monkeys[currentMonkey.Number] = currentMonkey
	}

	return monkeys
}

func day11_part1() {
	fmt.Println("DAY11: PART 1")
	monkeys := day11_import_monkeys()
	var monkeyKeys []int
	for _, monkey := range monkeys {
		monkeyKeys = append(monkeyKeys, monkey.Number)
	}

	sort.Ints(monkeyKeys)

	for round := 1; round <= 20; round++ {
		fmt.Printf("\n-------------------------------\nROUND %d\n", round)
		for _, monkeyNum := range monkeyKeys {
			monkey := monkeys[monkeyNum]
			fmt.Printf("  monkey %d\n", monkey.Number)
			for true {
				pending, target, newItem := monkey.InspectItem(true)
				if !pending {
					break
				}
				//fmt.Printf("Setting item %d to monkey %d\n", newItem, target)
				m := monkeys[target]
				m.Items = append(m.Items, newItem)
				monkeys[target] = m
			}

			monkeys[monkeyNum] = monkey
		}
	}

	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.Number, monkey.ItemsInspected)
	}
}
func day11_part2() {
	fmt.Println("DAY11: PART 2")
	monkeys := day11_import_monkeys()
	divisibleBy := int64(1)
	var monkeyKeys []int
	for _, monkey := range monkeys {
		monkeyKeys = append(monkeyKeys, monkey.Number)
		divisibleBy *= int64(monkey.TestDivisibleBy)
	}

	sort.Ints(monkeyKeys)

	for round := 1; round <= 10000; round++ {
		//fmt.Printf("\n-------------------------------\nROUND %d\n", round)
		for _, monkeyNum := range monkeyKeys {
			monkey := monkeys[monkeyNum]
			//fmt.Printf("  monkey %d\n", monkey.Number)
			for true {
				pending, target, newItem := monkey.InspectItem(false)
				if !pending {
					break
				}
				//fmt.Printf("Setting item %d to monkey %d\n", newItem, target)
				m := monkeys[target]
				m.Items = append(m.Items, newItem)
				monkeys[target] = m
			}

			monkeys[monkeyNum] = monkey
			//time.Sleep(time.Second * 2)
		}
		for m, _ := range monkeys {
			monkey := monkeys[m]
			monkey.Reduce(divisibleBy)
			monkeys[m] = monkey
		}
	}

	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.Number, monkey.ItemsInspected)
	}
}
