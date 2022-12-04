package day4

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
)

func part1(r io.Reader) {
	sum := 0
	var s, s1, s2, s3, s4, s5, s6 string
	var d1, d2, d3, d4 int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		s1, s2, _ = strings.Cut(s, ",")
		s3, s4, _ = strings.Cut(s1, "-")
		s5, s6, _ = strings.Cut(s2, "-")
		d1, _ = strconv.Atoi(s3)
		d2, _ = strconv.Atoi(s4)
		d3, _ = strconv.Atoi(s5)
		d4, _ = strconv.Atoi(s6)
		if d1 == d3 || d2 == d4 {
			sum++
			continue
		}
		if d3 > d1 {
			if d4 < d2 {
				sum++
			}
		} else {
			if d2 < d4 {
				sum++
			}
		}
	}
	//fmt.Println(sum)
}

func part2(r io.Reader) {
	sum := 0
	var s, s1, s2, s3, s4, s5, s6 string
	var d1, d2, d3, d4 int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		s1, s2, _ = strings.Cut(s, ",")
		s3, s4, _ = strings.Cut(s1, "-")
		s5, s6, _ = strings.Cut(s2, "-")
		d1, _ = strconv.Atoi(s3)
		d2, _ = strconv.Atoi(s4)
		d3, _ = strconv.Atoi(s5)
		d4, _ = strconv.Atoi(s6)
		if !(d3 > d2 || d1 > d4) {
			sum++
		}
	}
	//fmt.Println(sum)
}

func Variant1() {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
