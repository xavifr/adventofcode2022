package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day14_part1()
	day14_part2()
}

type BlockData int

const (
	Air BlockData = iota
	Sand
	Rock
	Generator
)

type Point struct {
	X int
	Y int
}

type SandCastle struct {
	W         int
	H         int
	Points    [][]BlockData
	Generator Point
	vX        int
	vY        int
	vW        int
	vH        int
}

func (sc *SandCastle) AddRock(X, Y, W, H int) {
	if X < 0 || Y < 0 || W < 0 || H < 0 || X+W > sc.W || Y+H > sc.H {
		return
	}

	for i := Y; i <= Y+H; i++ {
		for j := X; j <= X+W; j++ {
			sc.Points[i][j] = Rock
		}
	}
}

func (sc *SandCastle) Generate() Point {
	x := sc.Generator.X
	y := sc.Generator.Y
	for ; y < sc.H-1; y++ {
		if x >= sc.W || x < 0 {
			break
		}
		if sc.Points[y+1][x] == Air {
			continue
		} else if sc.Points[y+1][x-1] == Air {
			x -= 1
		} else if sc.Points[y+1][x+1] == Air {
			x += 1
		} else {
			break
		}
	}

	sc.Points[y][x] = Sand

	return Point{X: x, Y: y}
}

func (sc *SandCastle) Draw() {
	fmt.Println("\033[0:0H")
	sandColor := color.New(color.BgYellow)
	rockColor := color.New(color.FgHiWhite)
	generatorColor := color.New(color.FgHiGreen)

	//for y := 0; y < sc.H; y++ {
	for y := sc.vY; y < sc.vY+sc.vH; y++ {
		fmt.Println("")
	}

	fmt.Printf("\033[0:0H")

	//for y := 0; y < sc.H; y++ {
	//	for x := 0; x < sc.W; x++ {
	for y := sc.vY; y <= sc.vY+sc.vH; y++ {
		for x := sc.vX; x <= sc.vX+sc.vW; x++ {
			fmt.Printf("\033[%d;%dH", y-sc.vY+1, x-sc.vX)

			switch sc.Points[y][x] {
			case Sand:
				sandColor.Printf(" ")
			case Rock:
				rockColor.Printf("#")
			case Generator:
				generatorColor.Printf("+")
			case Air:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\033[%d;%dH", sc.vH+2, 0)

}

func NewSandCastle(W, H, X, Y, vX, vY, vW, vH int) SandCastle {
	sc := SandCastle{W: W, H: H, vX: vX, vY: vY, vW: vW, vH: vH, Generator: Point{X: X, Y: Y}}
	sc.Points = make([][]BlockData, H)
	for i, _ := range sc.Points {
		sc.Points[i] = make([]BlockData, W)
	}

	sc.Points[Y][X] = Generator

	return sc
}

func day14_part1() {
	fmt.Println("DAY14: PART 1")
	sandCastle := day14_create_sand_castle()

	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	sandSettled := 0
	exit := false
	for !exit {
		lastSand := sandCastle.Generate()
		//sandCastle.Draw()
		//time.Sleep(time.Second)

		if lastSand.Y == sandCastle.H-1 {
			fmt.Printf("Sand touched the grass after %d grains at point %d:%d\n", sandSettled, lastSand.X, lastSand.Y)
			break
		}

		sandSettled++
	}
	//sandCastle.Draw()

}
func day14_part2() {
	fmt.Println("DAY14: PART 1")
	sandCastle := day14_create_sand_castle()

	lastRock := 0
	for y := 0; y < sandCastle.H; y++ {
		for x := 0; x < sandCastle.W; x++ {
			if sandCastle.Points[y][x] == Rock && y > lastRock {
				lastRock = y
				break
			}
		}
	}

	sandCastle.AddRock(0, lastRock+2, sandCastle.W-1, 0)
	sandCastle.vH = lastRock + 3 - 12
	sandCastle.vY = 12

	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	sandSettled := 0
	exit := false
	for !exit {
		lastSand := sandCastle.Generate()
		//sandCastle.Draw()
		//time.Sleep(time.Second)
		sandSettled++

		if lastSand.Y == sandCastle.Generator.Y && lastSand.X == sandCastle.Generator.X {
			fmt.Printf("Sand touched the generator after %d grains at point %d:%d\n", sandSettled, lastSand.X, lastSand.Y)
			break
		}

	}
	//sandCastle.Draw()

}

func day14_create_sand_castle() SandCastle {
	sandCastle := SandCastle{}
	file, err := os.Open("day14.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return sandCastle
	}

	defer file.Close()

	// MAIN
	//sandCastle = NewSandCastle(1000, 500, 500, 0, 460, 0, 80, 150)
	sandCastle = NewSandCastle(1100, 500, 500, 0, 0, 0, 1099, 499)
	// DEMO
	//sandCastle = NewSandCastle(1000, 500, 500, 0, 480, 0, 40, 30)

	reRock, _ := regexp.Compile(`^(\d+),(\d+)(?: \-\> (\d+),(\d+))`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		for {
			if !reRock.MatchString(input) {
				break
			}

			match := reRock.FindStringSubmatch(input)
			X, _ := strconv.Atoi(match[1])
			Y, _ := strconv.Atoi(match[2])
			W, _ := strconv.Atoi(match[3])
			H, _ := strconv.Atoi(match[4])

			sandCastle.AddRock(
				int(math.Min(float64(X), float64(W))),
				int(math.Min(float64(Y), float64(H))),
				int(math.Abs(float64(W-X))),
				int(math.Abs(float64(H-Y))))

			input = input[len(match[1])+len(match[2])+5:]
		}
	}

	return sandCastle
}
