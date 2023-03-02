package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	"os/exec"
	"time"
)

func main() {
	myMap, initialNode := day12_import_map()

	b := make([]byte, 1)

	fmt.Println("DAY12: PART 1")
	found, node := day12_part1(myMap, initialNode, false)
	if found {
		fmt.Printf("%+v\n", node)
	} else {
		fmt.Printf("  not found\n")
	}
	os.Stdin.Read(b)

	fmt.Println("DAY12: PART 2")

	for y := 0; y < myMap.H; y++ {
		for x := 0; x < myMap.W; x++ {
			if myMap.Map[y][x] == 'a' {
				myNewMap, _ := day12_import_map()
				found, node := day12_part1(myNewMap, myNewMap.Nodes[y][x], false)
				if found {
					fmt.Printf("Walking from %d:%d\n", y, x)
					fmt.Printf("%+v\n", node)
				}
			}
		}
	}

	os.Stdin.Read(b)

	fmt.Println("DAY12: Fancy")
	myMap, initialNode = day12_import_map()
	_, n := day12_part1(myMap, initialNode, true)
	fmt.Printf("%+v\n", n)

}

type Map struct {
	W     int
	H     int
	Nodes [][]Node
	Map   [][]rune
}

func (m Map) DrawPath() {
	m.Draw()
	endNode := Node{}
	for y := 0; y < m.H; y++ {
		for x := 0; x < m.W; x++ {
			if m.Nodes[y][x].End {
				endNode = m.Nodes[y][x]
			}
		}
	}

	pathColor := color.New(color.FgBlue)
	currentNode := endNode
	for {
		fmt.Printf("\033[%d;%dH", currentNode.Y+2, currentNode.X+2)
		pathColor.Printf("%c", m.Map[currentNode.ParentY][currentNode.ParentX])

		if currentNode.Start {
			break
		}

		currentNode = m.Nodes[currentNode.ParentY][currentNode.ParentX]
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("\033[%d;%dH", m.H+2, 0)
}
func (m Map) Draw() {

	fmt.Println("\033[0:0H")
	startColor := color.New(color.FgGreen)
	endColor := color.New(color.FgRed)
	nextColor := color.New(color.FgHiBlue)
	visitedColor := color.New(color.FgHiGreen)
	for y := 0; y < m.H; y++ {
		for x := 0; x < m.W; x++ {
			printFunc := fmt.Printf
			if m.Nodes[y][x].Visited {
				printFunc = visitedColor.Printf
			} else if m.Nodes[y][x].Score != 99999 {
				printFunc = nextColor.Printf
			}

			if m.Nodes[y][x].Start {
				startColor.Printf("S")
			} else if m.Nodes[y][x].End {
				endColor.Printf("E")
			} else {
				printFunc("%c", m.Map[y][x])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\033[%d;%dH", m.H+2, 0)

}

type Node struct {
	X       int
	Y       int
	Visited bool
	Score   int
	Start   bool
	End     bool
	ParentX int
	ParentY int
}

func day12_part1(myMap Map, initialNode Node, draw bool) (bool, Node) {
	var unvisitedNodes []Node
	var nextUnvisitedNodes []Node
	unvisitedNodes = []Node{}

	myMap.Nodes[initialNode.Y][initialNode.X].Score = 0
	endNode := Node{Score: 99999}
	unvisitedNodes = append(unvisitedNodes, initialNode)

	if draw {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	round := 0
	for true {
		//fmt.Printf("---------------------------\nROUND %d, unvisited %d\n", round, len(unvisitedNodes))

		if len(unvisitedNodes) == 0 {
			//fmt.Println("No unvisited left :(")
			break
		}

		nextUnvisitedNodes = make([]Node, 0)
		for _, node := range unvisitedNodes {
			myNode := myMap.Nodes[node.Y][node.X]
			myHeight := myMap.Map[node.Y][node.X]
			//fmt.Printf("AT NODE %d:%d WITH HEIGHT %c\n", node.Y, node.X, myHeight)

			neighbours := make([]Node, 0)
			if node.X > 0 {
				//if math.Abs(float64(myHeight-myMap.Map[node.Y][node.X-1])) <= 1 {
				if math.Abs(float64(myHeight-myMap.Map[node.Y][node.X-1])) <= 1 || myHeight-myMap.Map[node.Y][node.X-1] > 0 {
					neighbours = append(neighbours, myMap.Nodes[node.Y][node.X-1])
				}
			}
			if node.X < myMap.W-1 {
				if math.Abs(float64(myHeight-myMap.Map[node.Y][node.X+1])) <= 1 || myHeight-myMap.Map[node.Y][node.X+1] > 0 {
					neighbours = append(neighbours, myMap.Nodes[node.Y][node.X+1])
				}
			}
			if node.Y > 0 {
				if math.Abs(float64(myHeight-myMap.Map[node.Y-1][node.X])) <= 1 || myHeight-myMap.Map[node.Y-1][node.X] > 0 {
					neighbours = append(neighbours, myMap.Nodes[node.Y-1][node.X])
				}
			}
			if node.Y < myMap.H-1 {
				if math.Abs(float64(myHeight-myMap.Map[node.Y+1][node.X])) <= 1 || myHeight-myMap.Map[node.Y+1][node.X] > 0 {
					neighbours = append(neighbours, myMap.Nodes[node.Y+1][node.X])
				}
			}

			for _, neighbour := range neighbours {
				if neighbour.Visited {
					continue
				}
				//fmt.Printf("  neigh: %d:%d => %c\n", neighbour.Y, neighbour.X, myMap.Map[neighbour.Y][neighbour.X])

				if myNode.Score+1 < neighbour.Score {
					//fmt.Printf("    score improved from %d to %d\n", neighbour.Score, myNode.Score+1)
					myMap.Nodes[neighbour.Y][neighbour.X].Score = myNode.Score + 1
					myMap.Nodes[neighbour.Y][neighbour.X].ParentY = myNode.Y
					myMap.Nodes[neighbour.Y][neighbour.X].ParentX = myNode.X
				} else {
					//fmt.Printf("    actual score was %d <=> %d\n", myNode.Score, neighbour.Score)
				}

				skipNode := false
				for _, nun := range nextUnvisitedNodes {
					if nun.X == neighbour.X && nun.Y == neighbour.Y {
						skipNode = true
					}
				}

				if !skipNode {
					nextUnvisitedNodes = append(nextUnvisitedNodes, myMap.Nodes[neighbour.Y][neighbour.X])
				}
			}

			myMap.Nodes[node.Y][node.X].Visited = true

			if myNode.End {
				endNode = myNode
				break
			}
			unvisitedNodes = nextUnvisitedNodes
			if draw {
				myMap.Draw()
			}
			//time.Sleep(time.Millisecond / 10)
		}

		if endNode.End {
			//fmt.Println("END FOUND")
			//fmt.Printf("%+v\n", endNode)
			if draw {
				myMap.DrawPath()
			}
			return true, endNode
		}
		round++
	}

	return false, endNode
}

func day12_import_map() (Map, Node) {
	myMap := Map{}
	initialNode := Node{}

	file, err := os.Open("day12.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return myMap, initialNode
	}

	defer file.Close()

	scanner := bufio.NewReader(file)
	x := 0
	y := 0

	myMap.Map = [][]rune{}
	myMap.Map = append(myMap.Map, []rune{})
	myMap.Nodes = [][]Node{}
	myMap.Nodes = append(myMap.Nodes, []Node{})
	for {
		input, _, err := scanner.ReadRune()
		if err != nil {
			break
		}

		if input == '\n' {
			y++
			x = 0
			myMap.Map = append(myMap.Map, []rune{})
			myMap.Nodes = append(myMap.Nodes, []Node{})
			continue
		}

		currentNode := Node{Visited: false, Score: 99999, X: x, Y: y}
		if input == 'S' {
			currentNode.Start = true
			currentNode.Score = 0
			initialNode = currentNode
			input = 'a'
		} else if input == 'E' {
			currentNode.End = true
			input = 'z'
		}

		myMap.Map[y] = append(myMap.Map[y], input)
		myMap.Nodes[y] = append(myMap.Nodes[y], currentNode)
		x++
	}

	if len(myMap.Map[len(myMap.Map)-1]) == 0 {
		myMap.Map = myMap.Map[0 : len(myMap.Map)-1]
	}
	if len(myMap.Nodes[len(myMap.Nodes)-1]) == 0 {
		myMap.Nodes = myMap.Nodes[0 : len(myMap.Nodes)-1]
	}

	myMap.H = len(myMap.Map)
	myMap.W = len(myMap.Map[0])

	return myMap, initialNode
}
