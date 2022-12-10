package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func part1(buf io.Reader) {
	var s1, s2 string
	var tx, ty, hx, hy, n, dx, dy int
	m := make(map[Point]struct{})
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		s1, s2, _ = strings.Cut(scanner.Text(), " ")
		n, _ = strconv.Atoi(s2)
		for i := 0; i < n; i++ {
			switch s1 {
			case "U":
				hy++
			case "R":
				hx++
			case "D":
				hy--
			case "L":
				hx--
			}
			dx = hx - tx
			dy = hy - ty

			if dx*dx+dy*dy <= 2 {
				continue
			}

			if dx > 1 {
				dx = 1
			}
			if dx < -1 {
				dx = -1
			}
			if dy > 1 {
				dy = 1
			}
			if dy < -1 {
				dy = -1
			}
			tx += dx
			ty += dy
			m[Point{tx, ty}] = struct{}{}
		}
	}
	fmt.Println(len(m) + 1)
}

func part2(buf io.Reader) {
	var s1, s2 string
	var tx, ty, hx, hy, n, dx, dy int
	rope := make([]Point, 10, 10)
	m := make(map[Point]struct{})
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		s1, s2, _ = strings.Cut(scanner.Text(), " ")
		n, _ = strconv.Atoi(s2)
		for i := 0; i < n; i++ {
			switch s1 {
			case "U":
				rope[0] = Point{rope[0].x, rope[0].y + 1}
			case "R":
				rope[0] = Point{rope[0].x + 1, rope[0].y}
			case "D":
				rope[0] = Point{rope[0].x, rope[0].y - 1}
			case "L":
				rope[0] = Point{rope[0].x - 1, rope[0].y}
			}
			for q := 1; q < 10; q++ {
				hx, hy = rope[q-1].x, rope[q-1].y
				tx, ty = rope[q].x, rope[q].y
				dx = hx - tx
				dy = hy - ty

				if dx*dx+dy*dy <= 2 {
					continue
				}

				if dx > 1 {
					dx = 1
				}
				if dx < -1 {
					dx = -1
				}
				if dy > 1 {
					dy = 1
				}
				if dy < -1 {
					dy = -1
				}
				rope[q] = Point{tx + dx, ty + dy}
				if q == 9 {
					m[Point{tx, ty}] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(m) + 1)
}

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
