package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Move struct {
	Direction string
	Length    int
}

type Wire struct {
	Moves []Move
}

type Segment struct {
	XStart, YStart int
	XEnd, YEnd     int
}

func NewSegment(xStart int, yStart int, xEnd int, yEnd int) Segment {
	if xStart < xEnd || yStart < yEnd {
		return Segment{xStart, yStart, xEnd, yEnd}
	} else if xEnd < xStart || yEnd < yStart {
		return Segment{xEnd, yEnd, xStart, yStart}
	} else {
		panic(fmt.Sprintf("Invalid segment %v,%v - %v,%v", xStart, yStart, xEnd, yEnd))
	}
}

func (s *Segment) IsVertical() bool {
	return s.XStart == s.XEnd
}

func (s *Segment) IsHorizontal() bool {
	return s.YStart == s.YEnd
}

func (s1 *Segment) Crosses(s2 Segment) (bool, Point) {
	if s1.IsVertical() && s2.IsHorizontal() {
		if s2.XStart <= s1.XStart && s1.XStart <= s2.XEnd &&
			s1.YStart <= s2.YStart && s2.YStart <= s1.YEnd {
			point := Point{s1.XStart, s2.YStart}
			fmt.Printf("s1: %v crosses s2: %v at %v\n", s1, s2, point)
			return point != Point{0,0}, point
		}
	} else if s1.IsHorizontal() && s2.IsVertical() {
		return s2.Crosses(*s1)
	}

	return false, Point{0, 0}
}

func FindIntersections(s1s []Segment, s2s []Segment) []Point {
	var res []Point

	for _, s1 := range s1s {
		for _, s2 := range s2s {
			if crosses, p := s1.Crosses(s2); crosses {
				res = append(res, p)
			}
		}
	}

	return res
}

func FindClosestIntersection(points []Point) Point {
	res := Point{0,0}
	minDist := math.Inf(+1)

	for _, p := range points {
		dist := math.Abs(float64(p.X)) + math.Abs(float64(p.Y))
		if dist < minDist {
			res = p
			minDist = dist
		}
	}

	return res
}

func (ant *Wire) getSegments() []Segment {
	var res []Segment

	x, y := 0, 0

	for _, move := range ant.Moves {
		xEnd := x
		yEnd := y
		switch move.Direction {
		case "R":
			xEnd += move.Length
		case "U":
			yEnd += move.Length
		case "D":
			yEnd -= move.Length
		case "L":
			xEnd -= move.Length
		}
		res = append(res, NewSegment(x, y, xEnd, yEnd))
		x = xEnd
		y = yEnd

	}

	return res
}

func scanAnt(scanner *bufio.Scanner) *Wire {
	r, _ := regexp.Compile("([RULD])([0-9]+)")

	s := scanner.Text()
	var moves []Move

	for _, i := range strings.Split(s, ",") {
		res := r.FindStringSubmatch(i)
		dir := res[1]
		len, _ := strconv.Atoi(res[2])
		moves = append(moves, Move{dir, len})
	}

	return &Wire{moves}
}

func main() {
	var input string = "-"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	var file *os.File
	if input == "-" {
		file = os.Stdin
	} else {
		_file, err := os.Open(input)
		if err != nil {
			panic(fmt.Sprintf("Could not open %v", input))
		}
		defer file.Close()
		file = _file
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		a1 := scanAnt(scanner)
		s1 := a1.getSegments()
		fmt.Println(s1)
		scanner.Scan()
		a2 := scanAnt(scanner)
		s2 := a2.getSegments()
		fmt.Println(s2)

		intersections := FindIntersections(s1, s2)
		fmt.Printf("Intersections: %v\n", intersections)
		fmt.Printf("Closest: %v\n", FindClosestIntersection(intersections))
		fmt.Println()
	}
}
