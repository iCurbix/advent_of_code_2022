package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func updateAround(terrain [][]uint8, steps [][]int, p point, step int) (bool, []point) {
	ter := terrain[p.y][p.x]
	var newPoints []point
	if ter == 'z' || ter == 'y' {
		if p.x-1 >= 0 && terrain[p.y][p.x-1] == 'E' {
			return true, nil
		}
		if p.x+1 < len(terrain[0]) && terrain[p.y][p.x+1] == 'E' {
			return true, nil
		}
		if p.y-1 >= 0 && terrain[p.y-1][p.x] == 'E' {
			return true, nil
		}
		if p.y+1 < len(terrain) && terrain[p.y+1][p.x] == 'E' {
			return true, nil
		}
	}
	if p.x-1 >= 0 && steps[p.y][p.x-1] == 0 && terrain[p.y][p.x-1]-ter+28 <= 29 {
		steps[p.y][p.x-1] = step
		newPoints = append(newPoints, point{x: p.x - 1, y: p.y})
	}
	if p.x+1 < len(terrain[0]) && steps[p.y][p.x+1] == 0 && terrain[p.y][p.x+1]-ter+28 <= 29 {
		steps[p.y][p.x+1] = step
		newPoints = append(newPoints, point{x: p.x + 1, y: p.y})
	}
	if p.y-1 >= 0 && steps[p.y-1][p.x] == 0 && terrain[p.y-1][p.x]-ter+28 <= 29 {
		steps[p.y-1][p.x] = step
		newPoints = append(newPoints, point{x: p.x, y: p.y - 1})
	}
	if p.y+1 < len(terrain) && steps[p.y+1][p.x] == 0 && terrain[p.y+1][p.x]-ter+28 <= 29 {
		steps[p.y+1][p.x] = step
		newPoints = append(newPoints, point{x: p.x, y: p.y + 1})
	}
	return false, newPoints
}

func updateAround2(terrain [][]uint8, steps [][]int, p point, step int) (bool, []point) {
	ter := terrain[p.y][p.x]
	var newPoints []point
	if ter == 'b' {
		if p.x-1 >= 0 && terrain[p.y][p.x-1] == 'a' {
			return true, nil
		}
		if p.x+1 < len(terrain[0]) && terrain[p.y][p.x+1] == 'a' {
			return true, nil
		}
		if p.y-1 >= 0 && terrain[p.y-1][p.x] == 'a' {
			return true, nil
		}
		if p.y+1 < len(terrain) && terrain[p.y+1][p.x] == 'a' {
			return true, nil
		}
	}
	if p.x-1 >= 0 && steps[p.y][p.x-1] == 0 && terrain[p.y][p.x-1]+1 >= ter {
		steps[p.y][p.x-1] = step
		newPoints = append(newPoints, point{x: p.x - 1, y: p.y})
	}
	if p.x+1 < len(terrain[0]) && steps[p.y][p.x+1] == 0 && terrain[p.y][p.x+1]+1 >= ter {
		steps[p.y][p.x+1] = step
		newPoints = append(newPoints, point{x: p.x + 1, y: p.y})
	}
	if p.y-1 >= 0 && steps[p.y-1][p.x] == 0 && terrain[p.y-1][p.x]+1 >= ter {
		steps[p.y-1][p.x] = step
		newPoints = append(newPoints, point{x: p.x, y: p.y - 1})
	}
	if p.y+1 < len(terrain) && steps[p.y+1][p.x] == 0 && terrain[p.y+1][p.x]+1 >= ter {
		steps[p.y+1][p.x] = step
		newPoints = append(newPoints, point{x: p.x, y: p.y + 1})
	}
	return false, newPoints
}

func do(terrain [][]uint8, start point, part int) {
	var updateFunc func(terrain [][]uint8, steps [][]int, p point, step int) (bool, []point)
	switch part {
	case 1:
		updateFunc = updateAround
	case 2:
		updateFunc = updateAround2
	}

	lx := len(terrain[0])
	ly := len(terrain)
	steps := make([][]int, ly, ly)
	for i := range steps {
		steps[i] = make([]int, lx, lx)
	}

	steps[start.y][start.x] = -1

	currPoints := make([]point, 1, 16)
	currPoints[0] = start
	i := 1
LOOP:
	for {
		var newCurrPoints []point
		for _, p := range currPoints {
			end, newPoints := updateFunc(terrain, steps, p, i)
			if end {
				break LOOP
			}
			newCurrPoints = append(newCurrPoints, newPoints...)
		}

		if newCurrPoints == nil {
			panic(fmt.Errorf("something bad happened"))
		}
		currPoints = newCurrPoints
		i++
	}
	fmt.Println(i)
}

func main() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var terrain [][]uint8
	scanner := bufio.NewScanner(f)
	var s string
	var start, end point
	i := 0
	for scanner.Scan() {
		s = scanner.Text()
		terrain = append(terrain, []uint8(s))
		if ind := strings.Index(s, "S"); ind >= 0 {
			start.x = ind
			start.y = i
		}
		if ind := strings.Index(s, "E"); ind >= 0 {
			end.x = ind
			end.y = i
		}
		i++
	}

	terrain[start.y][start.x] = 'a' - 1
	do(terrain, start, 1)
	do(terrain, end, 2)
}
