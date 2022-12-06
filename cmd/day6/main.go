package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func part1(r io.Reader) {
	a := make([]byte, 1)
	b := make([]byte, 4)
	for i := 0; i < 3; i++ {
		r.Read(a)
		b[i] = a[0]
	}
	i := 3
	for {
		r.Read(a)
		b[i%4] = a[0]
		for j := 0; j < 4; j++ {
			for q := j + 1; q < 4; q++ {
				if b[j] == b[q] {
					goto NOGOOD
				}
			}
		}
		break
	NOGOOD:
		i++
	}
	fmt.Println(i + 1)
}

func part2(r io.Reader) {
	a := make([]byte, 1)
	b := make([]byte, 14)
	for i := 0; i < 13; i++ {
		r.Read(a)
		b[i] = a[0]
	}
	i := 13
	for {
		r.Read(a)
		b[i%14] = a[0]
		for j := 0; j < 14; j++ {
			for q := j + 1; q < 14; q++ {
				if b[j] == b[q] {
					goto NOGOOD
				}
			}
		}
		break
	NOGOOD:
		i++
	}
	fmt.Println(i + 1)
}

func main() {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	io.ReadAll(tr)
	part2(buf)
}
