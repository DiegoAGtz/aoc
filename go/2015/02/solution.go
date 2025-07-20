package solution_2015

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(path string) []string {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	content := []string{}
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func partOne(dimensions []string) int {
	paperNeeded := 0
	for i := range len(dimensions) {
		l, w, h := getDimensions(dimensions[i])
		paperNeeded += perimeter(l, w, h) + area(l, w)
	}
	return paperNeeded
}

func partTwo(dimensions []string) int {
	ribbonNeeded := 0
	for i := range len(dimensions) {
		l, w, h := getDimensions(dimensions[i])
		ribbonNeeded += rectanglePerimeter(l, w) + volume(l, w, h)
	}
	return ribbonNeeded
}

func getDimensions(s string) (l, w, h int) {
	dims := strings.Split(s, "x")
	l, _ = strconv.Atoi(dims[0])
	w, _ = strconv.Atoi(dims[1])
	h, _ = strconv.Atoi(dims[2])
	l, w, h = sort(l, w, h)
	return
}

func perimeter(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func area(l, w int) int {
	return l * w
}

func rectanglePerimeter(l, w int) int {
	return 2*l + 2*w
}

func volume(l, w, h int) int {
	return l*w*h
}

func sort(l, w, h int) (a, b, c int) {
	if w < l {
		l += w
		w = l
		l -= w
	}
	if h < w {
		w += h
		h = w
		w -= h
	}
	return l, w, h
}
