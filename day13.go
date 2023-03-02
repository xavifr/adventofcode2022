package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type PacketType int

const (
	Int PacketType = iota
	List
)

type Packet struct {
	Name       string
	Message    int
	Type       PacketType
	Subpackets []Packet
}

func (p *Packet) Compare(p2 Packet) int {
	if p.Type == List && p2.Type != List {
		p2 = p2.ToList()
	} else if p2.Type == List && p.Type != List {
		*p = p.ToList()
	}

	if p.Type == Int && p2.Type == Int {
		if p.Message == p2.Message {
			return 0
		} else if p.Message < p2.Message {
			return 1
		}
		return -1
	}

	if (p.Subpackets == nil || len(p.Subpackets) == 0) && (p2.Subpackets == nil || len(p2.Subpackets) == 0) {
		return 0
	} else if p2.Subpackets == nil || len(p2.Subpackets) == 0 {
		return -1
	} else if p.Subpackets == nil || len(p.Subpackets) == 0 {
		return 1
	}

	i := 0
	for i = 0; i < int(math.Min(float64(len(p.Subpackets)), float64(len(p2.Subpackets)))); i++ {
		cmp := p.Subpackets[i].Compare(p2.Subpackets[i])
		if cmp == 0 {
			continue
		}

		return cmp
	}

	if len(p.Subpackets) > i {
		return -1
	} else if len(p2.Subpackets) > i {
		return 1
	}

	return 0
}

func (p *Packet) ToList() Packet {
	if p.Type == List {
		return *p
	}

	return Packet{Type: List, Subpackets: []Packet{*p}}
}

func day13_scan_packet(input string) Packet {
	p := Packet{Type: List, Subpackets: make([]Packet, 0)}
	if len(input) == 0 || input[0] != '[' || input[len(input)-1] != ']' || len(input) == 2 {
		return p
	}

	input = input[1 : len(input)-1]
	reInt, _ := regexp.Compile(`^(\d+)`)
	for {
		if len(input) == 0 {
			break
		} else if reInt.MatchString(input) { // int element
			match := reInt.FindStringSubmatch(input)
			intVal, _ := strconv.Atoi(match[1])
			p.Subpackets = append(p.Subpackets, Packet{Message: intVal, Type: Int})
			input = input[len(match[1]):]
		} else if input[0] == ',' { // next element
			input = input[1:]
		} else if input[0] == '[' {
			ct := 1
			i := 1
			for ; i < len(input) && ct != 0; i++ {
				if input[i] == '[' {
					ct++
				} else if input[i] == ']' {
					ct--
				}

			}

			p.Subpackets = append(p.Subpackets, day13_scan_packet(input[0:i]))

			input = input[i:]
		}

	}

	return p
}

func main() {
	day13_part1()
	day13_part2()
}

func day13_part1() {
	fmt.Println("DAY13: PART 1")

	file, err := os.Open("day13.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	packets := make([]Packet, 0)
	correctPackets := make([]int, 0)
	scanner := bufio.NewScanner(file)
	packetIDX := 1
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			// compare
			if packets[0].Compare(packets[1]) == 1 {
				correctPackets = append(correctPackets, packetIDX)
			}

			packets = make([]Packet, 0)
			packetIDX++
			continue
		}

		p := day13_scan_packet(scanner.Text())
		//fmt.Printf("%+v\n", p)
		packets = append(packets, p)
	}

	fmt.Printf("Correct packets are in position: %+v\n", correctPackets)
	res := 0
	for _, v := range correctPackets {
		res += v
	}
	fmt.Printf("Sum of indices is: %d\n", res)

}
func day13_part2() {
	fmt.Println("DAY13: PART 2")

	file, err := os.Open("day13.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return
	}

	defer file.Close()

	packets := make([]Packet, 0)
	p1 := day13_scan_packet("[[2]]")
	p1.Name = "M1"
	packets = append(packets, p1)
	p2 := day13_scan_packet("[[6]]")
	p2.Name = "M2"
	packets = append(packets, p2)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		p := day13_scan_packet(scanner.Text())
		//fmt.Printf("%+v\n", p)
		packets = append(packets, p)
	}

	sort.SliceStable(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) == 1
	})

	res := 1
	for i, p := range packets {
		if p.Name == "M1" {
			fmt.Printf("M1 found at %d\n", i+1)
			res *= i + 1
		} else if p.Name == "M2" {
			fmt.Printf("M2 found at %d\n", i+1)
			res *= i + 1
		}
	}

	fmt.Printf("Decoder key is %d\n", res)

}
