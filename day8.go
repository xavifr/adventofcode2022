package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	day8_part1()
	day8_part2()
}

func day8_part1() {
	fmt.Println("DAY8: PART 1")

	forest := day8_import_forest()

	visible := 0
	for x := 0; x < len(*forest); x++ {
		if x == 0 || x == len(*forest)-1 {
			visible += len((*forest)[x])
			//fmt.Printf("row %d is visible: %d\n", x, visible)
			continue
		}

		for y := 0; y < len((*forest)[x]); y++ {
			if y == 0 || y == len((*forest)[x])-1 {
				visible++
				//fmt.Printf("tree %d:%d (%d) is visible from sides: %d\n", x, y, (*forest)[x][y], visible)
				continue
			}

			// LEFT
			top_visible := true
			for z := x - 1; z >= 0; z-- {
				if (*forest)[x][y] <= (*forest)[z][y] {
					top_visible = false
					break
				}
			}
			if top_visible {
				visible++
				//fmt.Printf("tree %d:%d (%d) is visible from top: %d\n", x, y, (*forest)[x][y], visible)
				continue
			}

			// RIGHT
			down_visible := true
			for z := x + 1; z < len((*forest)[x]); z++ {
				if (*forest)[x][y] <= (*forest)[z][y] {
					down_visible = false
					break
				}
			}
			if down_visible {
				visible++
				//fmt.Printf("tree %d:%d (%d) is visible from down: %d\n", x, y, (*forest)[x][y], visible)
				continue
			}

			// TOP
			left_visible := true
			for z := y - 1; z >= 0; z-- {
				if (*forest)[x][y] <= (*forest)[x][z] {
					left_visible = false
					break
				}
			}
			if left_visible {
				visible++
				//fmt.Printf("tree %d:%d (%d) is visible from left: %d\n", x, y, (*forest)[x][y], visible)
				continue
			}

			// DOWN
			right_visible := true
			for z := y + 1; z < len(*forest); z++ {
				if (*forest)[x][y] <= (*forest)[x][z] {
					right_visible = false
					break
				}
			}
			if right_visible {
				visible++
				//fmt.Printf("tree %d:%d (%d) is visible from right: %d\n", x, y, (*forest)[x][y], visible)
				continue
			}

		}
	}

	fmt.Printf("There are %d visible trees\n", visible)
}
func day8_part2() {
	fmt.Println("DAY8: PART 2")

	forest := day8_import_forest()

	max_score := 0
	for x := 0; x < len(*forest); x++ {
		if x == 0 || x == len(*forest)-1 {
			continue
		}

		for y := 0; y < len((*forest)[x]); y++ {
			if y == 0 || y == len((*forest)[x])-1 {
				continue
			}

			// TOP
			//fmt.Printf("%d:%d Calc TOP\n", x, y)
			top_visible := 0
			for z := x - 1; z >= 0; z-- {
				//fmt.Printf("%d:%d   CMP %d <= %d\n", x, y, (*forest)[x][y], (*forest)[z][y])
				if (*forest)[x][y] <= (*forest)[z][y] {
					top_visible++
					//fmt.Printf("%d:%d     Break (%d)\n", x, y, top_visible)
					break
				}
				top_visible++
				//fmt.Printf("%d:%d     INC (%d)\n", x, y, top_visible)
			}

			// DOWN
			//fmt.Printf("%d:%d Calc DOWN\n", x, y)
			down_visible := 0
			for z := x + 1; z < len((*forest)[x]); z++ {
				//fmt.Printf("%d:%d   CMP %d <= %d\n", x, y, (*forest)[x][y], (*forest)[z][y])
				if (*forest)[x][y] <= (*forest)[z][y] {
					down_visible++
					//fmt.Printf("%d:%d     Break (%d)\n", x, y, down_visible)
					break
				}
				down_visible++
				//fmt.Printf("%d:%d     INC (%d)\n", x, y, down_visible)
			}

			// LEFT
			//fmt.Printf("%d:%d Calc LEFT\n", x, y)
			left_visible := 0
			for z := y - 1; z >= 0; z-- {
				//fmt.Printf("%d:%d   CMP %d <= %d\n", x, y, (*forest)[x][y], (*forest)[x][z])
				if (*forest)[x][y] <= (*forest)[x][z] {
					left_visible++
					//fmt.Printf("%d:%d     Break (%d)\n", x, y, left_visible)
					break
				}
				left_visible++
				//fmt.Printf("%d:%d     INC (%d)\n", x, y, left_visible)
			}

			// RIGHT
			//fmt.Printf("%d:%d Calc RIGHT\n", x, y)
			right_visible := 0
			for z := y + 1; z < len(*forest); z++ {
				//fmt.Printf("%d:%d   CMP %d <= %d\n", x, y, (*forest)[x][y], (*forest)[x][z])
				if (*forest)[x][y] <= (*forest)[x][z] {
					right_visible++
					//fmt.Printf("%d:%d     Break (%d)\n", x, y, right_visible)
					break
				}
				right_visible++
				//fmt.Printf("%d:%d     INC (%d)\n", x, y, right_visible)
			}

			if right_visible*left_visible*down_visible*top_visible > max_score {
				max_score = right_visible * left_visible * down_visible * top_visible
				//fmt.Printf("Improved max score at position %d:%d with %d:%d:%d:%d => %d\n", x, y, top_visible, down_visible, left_visible, right_visible, max_score)
			}
		}
	}

	fmt.Printf("Max score is %d \n", max_score)
}
func day8_import_forest() *[][]int {
	file, err := os.Open("day8.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return nil
	}

	defer file.Close()

	forest := new([][]int)

	scanner := bufio.NewReader(file)
	y := 0

	*forest = append(*forest, make([]int, 0))

	for {
		input, _, err := scanner.ReadRune()
		if err != nil {
			break
		}

		if input == '\n' {
			*forest = append(*forest, make([]int, 0))
			y++
			continue
		}

		(*forest)[y] = append((*forest)[y], int(input)-int('0'))
	}

	if len((*forest)[len(*forest)-1]) == 0 {
		*forest = (*forest)[0 : len(*forest)-1]
	}
	//fmt.Printf("%+v\n", forest)
	return forest
}
