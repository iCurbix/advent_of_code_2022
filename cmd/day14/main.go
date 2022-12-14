package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rang struct {
	down, up int
}

type point struct {
	x, y int
}

func getCave(r io.Reader) map[int][]*rang {
	var pts []string
	a := make(map[int][]*rang)
	var sx1, sx2, sy1, sy2 string
	var x1, x2, y1, y2 int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		pts = strings.Split(scanner.Text(), " -> ")
		for i := 1; i < len(pts); i++ {
			sx1, sy1, _ = strings.Cut(pts[i-1], ",")
			sx2, sy2, _ = strings.Cut(pts[i], ",")
			x1, _ = strconv.Atoi(sx1)
			y1, _ = strconv.Atoi(sy1)
			x2, _ = strconv.Atoi(sx2)
			y2, _ = strconv.Atoi(sy2)
			if x1 == x2 {
				if y2 < y1 {
					y1, y2 = y2, y1
				}
				a[x1] = append(a[x1], &rang{y2, y1})
			} else {
				q := 1
				if x1 > x2 {
					q = -1
				}
				if i != 1 {
					x1 += q
				}
				if i != len(pts)-1 {
					x2 -= q
				}

				if x1 > x2 {
					x1, x2 = x2, x1
				}

				for x := x1; x <= x2; x++ {
					a[x] = append(a[x], &rang{y1, y1})
				}
			}
		}
	}

	for _, v := range a {
		sort.Slice(v, func(l, r int) bool {
			return v[l].up < v[r].up
		})
	}

	return a
}

func part1(r io.Reader) {
	a := getCave(r)

	i := 1
	var pnt point
OUTER:
	for {
		pnt.x = 500
		pnt.y = 0
	INNER:
		for {
			var found bool
			for _, v := range a[pnt.x] {
				if v.up > pnt.y {
					pnt.y = v.up - 1
					found = true
					break
				}
			}
			if !found {
				break OUTER
			}

			found = false
			for _, v := range a[pnt.x-1] {
				if v.down > pnt.y {
					found = true
					if v.up <= pnt.y+1 {
						// to nie mozna na lewo
						break
					}
					pnt.x--
					pnt.y++
					continue INNER
				}
			}
			if !found {
				break OUTER
			}

			for _, v := range a[pnt.x+1] {
				if v.down > pnt.y {
					if v.up <= pnt.y+1 {
						// nie mozna w prawo
						break INNER
					}
					pnt.x++
					pnt.y++
					continue INNER
				}
			}
			break OUTER
		}

		for _, v := range a[pnt.x] {
			if v.up > pnt.y {
				v.up--
				break
			}
		}

		i++
	}

	fmt.Println(i - 1)
}

func part2(r io.Reader) {
	a := getCave(r)
	maxy := 0
	for _, v := range a {
		if ma := v[len(v)-1].down; ma > maxy {
			maxy = ma
		}
	}

	maxy += 2

	for i := 200; i < 800; i++ {
		a[i] = append(a[i], &rang{maxy, maxy})
	}

	i := 1
	var pnt point
OUTER:
	for {
		pnt.x = 500
		pnt.y = 0
	INNER:
		for {
			var found bool
			for _, v := range a[pnt.x] {
				if v.up > pnt.y {
					pnt.y = v.up - 1
					found = true
					break
				}
			}
			if !found {
				break OUTER
			}

			found = false
			for _, v := range a[pnt.x-1] {
				if v.down > pnt.y {
					found = true
					if v.up <= pnt.y+1 {
						// to nie mozna na lewo
						break
					}
					pnt.x--
					pnt.y++
					continue INNER
				}
			}
			if !found {
				break OUTER
			}

			for _, v := range a[pnt.x+1] {
				if v.down > pnt.y {
					if v.up <= pnt.y+1 {
						// nie mozna w prawo
						break INNER
					}
					pnt.x++
					pnt.y++
					continue INNER
				}
			}
			break OUTER
		}

		if pnt.x == 500 && pnt.y == 0 {
			break OUTER
		}

		for _, v := range a[pnt.x] {
			if v.up > pnt.y {
				v.up--
				break
			}
		}

		i++
	}

	fmt.Println(i)
}

func main() {
	f, err := os.Open("inputs/day14.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
