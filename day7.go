package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Type int64

const (
	Dir Type = iota
	File
)

type DIREntry struct {
	Name     string
	Type     Type
	Size     int64
	Children *[]DIREntry
	Parent   *DIREntry
}

func (de *DIREntry) GetSize() int64 {
	out := int64(0)
	if de.Children == nil || len(*de.Children) == 0 {
		return out
	}

	for _, d := range *de.Children {
		if d.Type == Dir {
			out += d.GetSize()
		} else {
			out += d.Size
		}
	}

	return out
}

func main() {
	day7_part1()
}

func day7_part1() {
	fmt.Println("DAY7: PART 1")
	root := day7_scan_commands()

	fmt.Printf("Total size is %d\n", root.GetSize())

	fmt.Printf("Recurse <100k is %d\n", day7_recurse_size_100k(*root))

	sizeMap := &map[string]int64{}
	day7_map_of_sizes(*root, sizeMap)

	fmt.Printf("%+v\n", sizeMap)

	totalSpace := int64(70000000)
	usedSpace := root.GetSize()
	desiredSpace := int64(30000000)
	needToFree := desiredSpace - (totalSpace - usedSpace)
	fmt.Printf("Need to free at least: %d\n", needToFree)

	minDelta := totalSpace
	minName := ""

	for sizeName, sizeVal := range *sizeMap {
		if sizeVal >= needToFree && sizeVal < minDelta {
			minDelta = sizeVal
			minName = sizeName
		}
	}

	fmt.Printf("Selected folder is %s with size of %d\n", minName, minDelta)

}

func day7_map_of_sizes(entry DIREntry, mapi *map[string]int64) {
	if entry.Type == File {
		return
	}

	(*mapi)[entry.Name] = entry.GetSize()

	for _, de := range *entry.Children {
		day7_map_of_sizes(de, mapi)
	}
}

func day7_recurse_size_100k(entry DIREntry) int64 {
	if entry.Type == File {
		return 0
	}

	acum := int64(0)
	if entry.GetSize() < 100000 {
		acum += entry.GetSize()
	}

	for _, de := range *entry.Children {
		if de.Type == Dir {
			acum += day7_recurse_size_100k(de)
		}
	}

	return acum
}

func day7_scan_commands() *DIREntry {
	file, err := os.Open("day7.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return nil
	}

	defer file.Close()

	rootDIREntry := DIREntry{Name: "/", Type: Dir, Children: &[]DIREntry{}}

	scanner := bufio.NewScanner(file)

	reCmdChangeDir, _ := regexp.Compile(`^\$ cd ([\w\./]+)$`)
	reCmdListDir, _ := regexp.Compile(`^\$ ls$`)
	reListDir, _ := regexp.Compile(`^dir ([\w\.]+)$`)
	reListFile, _ := regexp.Compile(`^(\d+)\s+([\w\.]+)$`)

	currentEntry := &rootDIREntry
	currentlyListing := false
	for scanner.Scan() {
		input := scanner.Text()

		if reCmdChangeDir.MatchString(input) {
			currentlyListing = false
			match := reCmdChangeDir.FindStringSubmatch(input)
			if match[1] == "/" {
				currentEntry = &rootDIREntry
			} else if match[1] == ".." {
				if currentEntry.Parent != nil {
					currentEntry = currentEntry.Parent
				} else {
					fmt.Printf("Cannot go folder UP from %s\n", currentEntry.Name)
				}
			} else {
				if currentEntry.Children != nil {
					found := false
					for _, dir := range *currentEntry.Children {
						if dir.Type == Dir && dir.Name == match[1] {
							fmt.Printf("CHDIR DIR %s\n", match[1])
							currentEntry = &dir
							found = true
							break
						}
					}
					if found {
						continue
					}
				}

				fmt.Printf("ADD DIR %s\n", match[1])
				dir := DIREntry{Name: match[1], Type: Dir, Children: &[]DIREntry{}, Parent: currentEntry}
				*currentEntry.Children = append(*currentEntry.Children, dir)
				currentEntry = &dir
			}
		} else if reCmdListDir.MatchString(input) {
			fmt.Printf("START LS\n")

			currentlyListing = true
		} else if currentlyListing && reListDir.MatchString(input) {
			match := reListDir.FindStringSubmatch(input)
			if currentEntry.Children != nil {
				found := false
				for _, dir := range *currentEntry.Children {
					if dir.Type == Dir && dir.Name == match[1] {
						fmt.Printf("RELOCATED DIR %s\n", match[1])
						found = true
						break
					}
				}

				if found {
					continue
				}
			}

			fmt.Printf("ADD DIR %s\n", match[1])
			dir := DIREntry{Name: match[1], Type: Dir, Children: &[]DIREntry{}, Parent: currentEntry}
			*currentEntry.Children = append(*currentEntry.Children, dir)
		} else if currentlyListing && reListFile.MatchString(input) {
			match := reListFile.FindStringSubmatch(input)
			size, _ := strconv.Atoi(match[1])

			if currentEntry.Children != nil {
				found := false
				for _, file := range *currentEntry.Children {
					if file.Type == File && file.Name == match[2] {
						file.Size = int64(size)
						fmt.Printf("RELOCATED FILE %s\n", match[2])
						found = true
						break
					}
				}

				if found {
					continue
				}
			}

			fmt.Printf("ADD FILE %s\n", match[2])
			*currentEntry.Children = append(*currentEntry.Children, DIREntry{Name: match[2], Type: File, Size: int64(size), Children: &[]DIREntry{}, Parent: currentEntry})
		}
	}

	return &rootDIREntry
}
