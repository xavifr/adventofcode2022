package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	day6_part1()
	day6_part2()
}

type FIFO []rune

func (s *FIFO) IsEmpty() bool {
	return len(*s) == 0
}

func (s *FIFO) Push(el rune) {
	*s = append(*s, el)
}

func (s *FIFO) Pop() (rune, bool) {
	if s.IsEmpty() {
		return ' ', false
	}

	el := (*s)[0]
	*s = (*s)[1:]

	return el, true
}

func (s *FIFO) Head() (rune, bool) {
	if s.IsEmpty() {
		return ' ', false
	}

	el := (*s)[len(*s)-1]

	return el, true
}

type Message struct {
	message   string
	marker    FIFO
	MarkerLen int
}

func (m *Message) Len() int {
	return len(m.message) + len(m.marker)
}

func (m *Message) ValidMarker() bool {
	if len(m.marker) != m.MarkerLen {
		return false
	}

	for i := 0; i < m.MarkerLen; i++ {
		for j := i + 1; j < m.MarkerLen; j++ {
			if m.marker[i] == m.marker[j] {
				return false
			}
		}
	}

	return true
}

func (m *Message) AppendRune(input rune) {
	m.marker.Push(input)
	if len(m.marker) <= m.MarkerLen {
		return
	}

	el, _ := m.marker.Pop()
	m.message = fmt.Sprintf("%s%c", m.message, el)
}

func day6_part1() {
	fmt.Println("DAY6: PART 1")
	file, err := os.Open("day6.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewReader(file)
	msg := Message{MarkerLen: 4}

	for {
		inputChar, _, e := scanner.ReadRune()
		if e != nil {
			break
		}

		msg.AppendRune(inputChar)
		if msg.ValidMarker() {
			fmt.Printf("MESSAGE LEN IS: %d\n", msg.Len())
			break
		}
	}

}
func day6_part2() {
	fmt.Println("DAY6: PART 2")
	file, err := os.Open("day6.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewReader(file)
	msg := Message{MarkerLen: 14}

	for {
		inputChar, _, e := scanner.ReadRune()
		if e != nil {
			break
		}

		msg.AppendRune(inputChar)
		if msg.ValidMarker() {
			fmt.Printf("MESSAGE LEN IS: %d\n", msg.Len())
			break
		}
	}

}
