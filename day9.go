package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func main() {
	b := make([]byte, 1)
	day9_part1()
	os.Stdin.Read(b)
	day9_part2()
	os.Stdin.Read(b)
	day9_part22()
	os.Stdin.Read(b)
}

type Snake struct {
	body          []Position
	len           int
	tailPostition map[string]int
}

func NewSnake(len int) Snake {
	s := Snake{len: len, body: make([]Position, len), tailPostition: map[string]int{}}
	s.tailPostition[s.body[0].Hash()] = 1

	return s
}

func (s *Snake) Move(dir rune, steps int) {
	fmt.Printf("MOVE %c LEN %d\n", dir, steps)
	for ct := 0; ct < steps; ct++ {
		if dir == 'R' {
			s.body[0].X += 1
		} else if dir == 'U' {
			s.body[0].Y -= 1
		} else if dir == 'L' {
			s.body[0].X -= 1
		} else if dir == 'D' {
			s.body[0].Y += 1
		}

		for i := 1; i < s.len; i++ {
			if !s.body[i].Adjacent(s.body[i-1]) {
				s.body[i].Follow(s.body[i-1])
			} else {
				break
			}
		}

		hash := s.body[s.len-1].Hash()
		if _, ok := s.tailPostition[hash]; !ok {
			s.tailPostition[hash] = 1
		} else {
			s.tailPostition[hash]++
		}
		s.Draw()

	}
	//	s.Draw()

	//	fmt.Printf("%+v\n", s.body)
}

func (s *Snake) Draw() {
	center := s.body[0]
	fmt.Println("\033[0:0H")
	for j := -10; j <= 10; j++ {
		for i := -40; i <= 40; i++ {
			drawn := false
			for pn, p := range s.body {
				if center.X+int64(i) == p.X && center.Y+int64(j) == p.Y {
					fmt.Printf("%d", pn)
					drawn = true
					break
				}
			}
			if !drawn {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	time.Sleep(time.Millisecond * 20)
}
func (s *Snake) GetTail() int {
	return len(s.tailPostition)
}

type Position struct {
	X int64
	Y int64
}

func (p *Position) Adjacent(o Position) bool {
	if p.X == o.X && p.Y == o.Y {
		return true
	}

	if p.X == o.X && math.Abs(float64(p.Y-o.Y)) == 1 {
		return true
	}
	if p.Y == o.Y && math.Abs(float64(p.X-o.X)) == 1 {
		return true
	}
	if math.Abs(float64(p.Y-o.Y)) == 1 && math.Abs(float64(p.X-o.X)) == 1 {
		return true
	}

	return false
}

func (p *Position) Follow(o Position) {
	if p.X == o.X { // moved Y
		if p.Y < o.Y {
			p.Y = o.Y - 1
		} else {
			p.Y = o.Y + 1
		}
	} else if p.Y == o.Y { // moved X
		if p.X < o.X {
			p.X = o.X - 1
		} else {
			p.X = o.X + 1
		}
	} else if math.Abs(float64(p.Y-o.Y)) == 2 && math.Abs(float64(p.X-o.X)) == 1 { // diagonal simple Y
		p.X = o.X
		if p.Y < o.Y {
			p.Y = o.Y - 1
		} else {
			p.Y = o.Y + 1
		}
	} else if math.Abs(float64(p.Y-o.Y)) == 1 && math.Abs(float64(p.X-o.X)) == 2 { // diagonal simple Y
		p.Y = o.Y
		if p.X < o.X {
			p.X = o.X - 1
		} else {
			p.X = o.X + 1
		}
	} else if math.Abs(float64(p.Y-o.Y)) == 2 && math.Abs(float64(p.X-o.X)) == 2 { // diagonal doble
		if p.Y < o.Y {
			p.Y = o.Y - 1
		} else {
			p.Y = o.Y + 1
		}
		if p.X < o.X {
			p.X = o.X - 1
		} else {
			p.X = o.X + 1
		}
	}

}

func (p *Position) Hash() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("X:%d Y:%d", p.X, p.Y))))
}

func day9_part1() {
	fmt.Println("DAY9: PART 1")

	file, err := os.Open("day9.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	head := Position{X: 500, Y: 500}
	tail := Position{X: 500, Y: 500}
	posHash := map[string]int{}
	posHash[tail.Hash()] = 1

	reMove, _ := regexp.Compile(`^([RULD]) (\d+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !reMove.MatchString(scanner.Text()) {
			continue
		}

		match := reMove.FindStringSubmatch(scanner.Text())
		moveCount, _ := strconv.Atoi(match[2])

		for i := 0; i < moveCount; i++ {
			if match[1] == "R" {
				head.X += 1
			} else if match[1] == "U" {
				head.Y -= 1
			} else if match[1] == "L" {
				head.X -= 1
			} else if match[1] == "D" {
				head.Y += 1
			}

			if !head.Adjacent(tail) {
				tail.Follow(head)
				hash := tail.Hash()
				if _, ok := posHash[hash]; !ok {
					posHash[hash] = 1
				} else {
					posHash[hash]++
				}
			}
		}
	}

	fmt.Printf("Tail has visited %d positions\n", len(posHash))
}

func day9_part2() {
	fmt.Println("DAY9: PART 2")

	file, err := os.Open("day9.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	snake := make([]Position, 10)

	posTail := map[string]int{}
	posTail[snake[len(snake)-1].Hash()] = 1

	reMove, _ := regexp.Compile(`^([RULD]) (\d+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !reMove.MatchString(scanner.Text()) {
			continue
		}

		match := reMove.FindStringSubmatch(scanner.Text())
		moveCount, _ := strconv.Atoi(match[2])

		for i := 0; i < moveCount; i++ {
			if match[1] == "R" {
				snake[0].X += 1
			} else if match[1] == "U" {
				snake[0].Y -= 1
			} else if match[1] == "L" {
				snake[0].X -= 1
			} else if match[1] == "D" {
				snake[0].Y += 1
			}

			for i := 1; i < len(snake); i++ {
				if !snake[i].Adjacent(snake[i-1]) {
					snake[i].Follow(snake[i-1])
				} else {
					break
				}
			}

			hash := snake[len(snake)-1].Hash()
			if _, ok := posTail[hash]; !ok {
				posTail[hash] = 1
			} else {
				posTail[hash]++
			}
		}
	}

	fmt.Printf("Tail has visited %d positions\n", len(posTail))

}
func day9_part22() {
	fmt.Println("DAY9: PART 2.1")

	file, err := os.Open("day9.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	snake := NewSnake(10)

	reMove, _ := regexp.Compile(`^([RULD]) (\d+)$`)

	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !reMove.MatchString(scanner.Text()) {
			continue
		}

		match := reMove.FindStringSubmatch(scanner.Text())
		moveCount, _ := strconv.Atoi(match[2])
		snake.Move(rune(match[1][0]), moveCount)

	}

	fmt.Printf("Tail has visited %d positions\n", snake.GetTail())

}
