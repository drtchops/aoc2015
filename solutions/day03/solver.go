package day03

import (
	"fmt"
	"strings"
)

type Solver struct {
	input string
}

func New(input string) *Solver {
	return &Solver{input: input}
}

// --- Day 3: Perfectly Spherical Houses in a Vacuum ---
// Santa is delivering presents to an infinite two-dimensional grid of houses.

// He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

// However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

// For example:

// > delivers presents to 2 houses: one at the starting location, and one to the east.
// ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
// ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

type point struct {
	X, Y int64
}

func move(pos point, dir string) point {
	if dir == "^" {
		return point{pos.X, pos.Y - 1}
	}
	if dir == "v" {
		return point{pos.X, pos.Y + 1}
	}
	if dir == ">" {
		return point{pos.X + 1, pos.Y}
	}
	if dir == "<" {
		return point{pos.X - 1, pos.Y}
	}
	return pos
}

func (s *Solver) SolveA() string {
	moves := strings.Split(s.input, "")
	pos := point{0, 0}
	visited := map[point]int64{pos: 1}

	for _, dir := range moves {
		pos = move(pos, dir)
		visited[pos] += 1
	}

	return fmt.Sprint(len(visited))
}

func (s *Solver) SolveB() string {
	moves := strings.Split(s.input, "")
	santaPos := point{0, 0}
	roboPos := point{0, 0}
	visited := map[point]int64{santaPos: 2}
	tick := 0

	for _, dir := range moves {
		if tick%2 == 0 {
			santaPos = move(santaPos, dir)
			visited[santaPos] += 1
		} else {
			roboPos = move(roboPos, dir)
			visited[roboPos] += 1
		}
		tick++
	}

	return fmt.Sprint(len(visited))
}
