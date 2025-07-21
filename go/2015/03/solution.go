package solutions_2015_03

import "fmt"

func PartOne(instructions string) int {
	n := len(instructions)
	m := make(map[string]int)
	m["0-0"]++
	x, y := 0, 0

	for i := range n {
		a, b := move(instructions[i])
		x += a
		y += b
		m[fmt.Sprintf("%v-%v", x, y)]++
	}

	return len(m)
}

func PartTwo(instructions string) int {
	n := len(instructions)
	m := make(map[string]int)
	m["0-0"]++
	x, y := 0, 0

	for i := 0; i < n; i += 2 {
		a, b := move(instructions[i])
		x += a
		y += b
		m[fmt.Sprintf("%v-%v", x, y)]++
	}

	x, y = 0, 0
	for i := 1; i < n; i += 2 {
		a, b := move(instructions[i])
		x += a
		y += b
		m[fmt.Sprintf("%v-%v", x, y)]++
	}

	return len(m)
}

func move(b byte) (x, y int) {
	x, y = 0, 0
	switch b {
	case '^':
		y++
	case '>':
		x++
	case 'v':
		y -= 1
	case '<':
		x -= 1
	}
	return
}
