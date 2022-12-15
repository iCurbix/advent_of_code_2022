package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Sensor struct {
	sx, sy, bx, by, dist int
}

const (
	ROW = 2000000
	MIN = 0
	MAX = 4000000
)

func main() {
	f, err := os.Open("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`-?\d+`)
	var sensors []*Sensor
	var sx, sy, bx, by, dx, dy, dist, minx, maxx int
	for scanner.Scan() {
		all := re.FindAllString(scanner.Text(), -1)
		sx, _ = strconv.Atoi(all[0])
		sy, _ = strconv.Atoi(all[1])
		bx, _ = strconv.Atoi(all[2])
		by, _ = strconv.Atoi(all[3])
		dx = sx - bx
		if dx < 0 {
			dx *= -1
		}
		dy = sy - by
		if dy < 0 {
			dy *= -1
		}
		dist = dx + dy
		if sx-dist < minx {
			minx = sx - dist
		}
		if sx+dist > maxx {
			maxx = sx + maxx
		}
		sensors = append(sensors, &Sensor{sx, sy, bx, by, dx + dy})
	}

	row := make([]bool, maxx-minx, maxx-minx)

	for _, s := range sensors {
		dist = s.sy - ROW
		if dist < 0 {
			dist *= -1
		}
		dif := s.dist - dist
		if dif < 0 {
			continue
		}
		for xx := s.sx - dif; xx < s.sx+dif+1; xx++ {
			row[xx-minx] = true
		}
	}

	for _, s := range sensors {
		if s.by == ROW {
			row[s.bx-minx] = false
		}
	}

	sum := 0
	for _, v := range row {
		if v {
			sum++
		}
	}

	fmt.Println(sum)
	//
	// PART 2
	//

OUTER:
	for y := MIN; y <= MAX; y++ {
		fmt.Println(y)
		row := make([]bool, maxx-minx, maxx-minx)
		for _, s := range sensors {
			dist = s.sy - y
			if dist < 0 {
				dist *= -1
			}
			dif := s.dist - dist
			if dif < 0 {
				continue
			}
			for xx := s.sx - dif; xx < s.sx+dif+1; xx++ {
				row[xx-minx] = true
			}
		}
		//fmt.Println(y, row[MIN-minx:MAX+1-minx])
		for i := MIN; i <= MAX; i++ {
			if !row[i-minx] {
				fmt.Println(i*4000000 + y)
				break OUTER
			}
		}
	}

}
