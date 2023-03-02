package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day15_part1()
	day15_part2()
}

type Point15 struct {
	X int
	Y int
}

type Sensor struct {
	Sensor Point15
	Beacon Point15
}

func (s *Sensor) DistanceFromSensor(p Point15) int {
	return int(math.Abs(float64(s.Sensor.X-p.X)) + math.Abs(float64(s.Sensor.Y-p.Y)))
}

func (s *Sensor) IsPointCoveredByBeacon(p Point15) bool {
	return s.DistanceFromSensor(p) <= s.DistanceFromSensor(s.Beacon)
}

func day15_part1() {
	fmt.Println("DAY15: PART 1")
	sensors := day15_read_sensors_data()

	minX := 99999999999
	maxX := 0
	for _, s := range sensors {
		if s.Sensor.X < minX {
			minX = s.Sensor.X
		}
		if (s.Sensor.X - s.DistanceFromSensor(s.Beacon)) < minX {
			minX = s.Sensor.X - s.DistanceFromSensor(s.Beacon)
		}

		if s.Sensor.X > maxX {
			maxX = s.Sensor.X
		}
		if (s.Sensor.X + s.DistanceFromSensor(s.Beacon)) > maxX {
			maxX = s.Sensor.X + s.DistanceFromSensor(s.Beacon)
		}
	}

	coveredPoints := make([]Point15, 0)
	y := 2000000
	for x := minX; x < maxX; x++ {
		covered := false
		point := Point15{x, y}
		for _, sensor := range sensors {
			if sensor.Beacon == point {
				continue
			}
			if sensor.IsPointCoveredByBeacon(point) {
				covered = true
				break
			}
		}

		if covered {
			coveredPoints = append(coveredPoints, point)
		}
	}

	fmt.Printf("Covered points at row %d ar %d\n", y, len(coveredPoints))
}
func day15_part2() {
	fmt.Println("DAY15: PART 2")
	sensors := day15_read_sensors_data()

	minX := 0
	maxX := 4000000
	minY := 0
	maxY := 4000000

	freePoints := make([]Point15, 0)

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			covered := false
			point := Point15{x, y}
			for _, sensor := range sensors {
				if sensor.Beacon == point || sensor.IsPointCoveredByBeacon(point) {
					covered = true
					//fmt.Printf("%d:%d - DFB: %d, DFP: %d, Jump %d\n", point.X, point.Y, sensor.DistanceFromSensor(sensor.Beacon), sensor.DistanceFromSensor(point), sensor.DistanceFromSensor(sensor.Beacon)-sensor.DistanceFromSensor(point))
					x += sensor.DistanceFromSensor(sensor.Beacon) - sensor.DistanceFromSensor(point)
					break
				}
			}

			if !covered {
				freePoints = append(freePoints, point)
				break
			}
		}

		if len(freePoints) > 0 {
			break
		}
	}

	if len(freePoints) == 1 {
		fmt.Printf("Distress beacon is at %d:%d with frequency %d\n", freePoints[0].X, freePoints[0].Y, freePoints[0].X*4000000+freePoints[0].Y)
	} else {
		fmt.Printf("Free points len is %d\n", len(freePoints))
	}
}

func day15_read_sensors_data() []Sensor {
	sensors := make([]Sensor, 0)
	file, err := os.Open("day15.input")
	if err != nil {
		fmt.Printf("Error reading input! %s\n", err.Error())
		return sensors
	}

	defer file.Close()

	reSensorBeacon, _ := regexp.Compile(`^Sensor at x=([\-\d]+), y=([\-\d]+): closest beacon is at x=([\-\d]+), y=([\-\d]+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !reSensorBeacon.MatchString(scanner.Text()) {
			continue
		}

		match := reSensorBeacon.FindStringSubmatch(scanner.Text())
		sensorX, _ := strconv.Atoi(match[1])
		sensorY, _ := strconv.Atoi(match[2])
		beaconX, _ := strconv.Atoi(match[3])
		beaconY, _ := strconv.Atoi(match[4])
		sensors = append(sensors, Sensor{Sensor: Point15{sensorX, sensorY}, Beacon: Point15{beaconX, beaconY}})
	}

	return sensors
}
