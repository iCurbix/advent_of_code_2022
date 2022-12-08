package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction uint

const (
	UP    direction = 1
	RIGHT direction = 1 << 1
	DOWN  direction = 1 << 2
	LEFT  direction = 1 << 3
)

func setBit(n direction, pos direction) direction {
	n |= pos
	return n
}

func clearBit(n direction, pos direction) direction {
	mask := ^pos
	n &= mask
	return n
}

func hasBit(n direction, pos direction) bool {
	val := n & pos
	return val > 0
}

func insertRow(row []int32, vis []direction, s string) {
	for i, v := range s {
		row[i] = v - 48
		vis[i] = 15
	}
}

func main() {
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	l := len(scanner.Text())
	h := make([][]int32, l, l)
	h[0] = make([]int32, l, l)
	vis := make([][]direction, l, l)
	vis[0] = make([]direction, l, l)

	insertRow(h[0], vis[0], scanner.Text())

	for i := 1; i < l; i++ {
		scanner.Scan()
		scanner.Text()
		h[i] = make([]int32, l, l)
		vis[i] = make([]direction, l, l)
		insertRow(h[i], vis[i], scanner.Text())
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if hasBit(vis[i][j], UP) {
				for q := i + 1; q < l; q++ {
					if h[i][j] >= h[q][j] {
						vis[q][j] = clearBit(vis[q][j], UP)
					}
				}
			}
			if hasBit(vis[i][j], LEFT) {
				for q := j + 1; q < l; q++ {
					if h[i][j] >= h[i][q] {
						vis[i][q] = clearBit(vis[i][q], LEFT)
					}
				}
			}
		}
	}

	for i := l - 1; i >= 0; i-- {
		for j := l - 1; j >= 0; j-- {
			if hasBit(vis[i][j], DOWN) {
				for q := i - 1; q >= 0; q-- {
					if h[i][j] >= h[q][j] {
						vis[q][j] = clearBit(vis[q][j], DOWN)
					}
				}
			}
			if hasBit(vis[i][j], RIGHT) {
				for q := j - 1; q >= 0; q-- {
					if h[i][j] >= h[i][q] {
						vis[i][q] = clearBit(vis[i][q], RIGHT)
					}
				}
			}
		}
	}

	sum := 0

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if vis[i][j] > 0 {
				sum++
			}
		}
	}
	fmt.Println(sum)

	maxScore := 0
	var u, r, d, le int

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			u, r, d, le = 0, 0, 0, 0
			// UP
			for q := i - 1; q >= 0; q-- {
				u++
				if h[q][j] >= h[i][j] {
					break
				}
			}
			// RIGHT
			for q := j + 1; q < l; q++ {
				r++
				if h[i][q] >= h[i][j] {
					break
				}
			}
			// DOWN
			for q := i + 1; q < l; q++ {
				d++
				if h[q][j] >= h[i][j] {
					break
				}
			}
			// LEFT
			for q := j - 1; q >= 0; q-- {
				le++
				if h[i][q] >= h[i][j] {
					break
				}
			}

			if u*r*d*le > maxScore {
				maxScore = u * r * d * le
			}
		}
	}
	fmt.Println(maxScore)
}
