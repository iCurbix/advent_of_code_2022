package day3

import (
	"advent_of_code/pkg/buffer"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func part1(buf io.Reader) {
	sum := 0
	l := 0
	i := 0
	s := ""
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		s = scanner.Text()
		l = len(s) / 2
		i = strings.IndexAny(s[:l], s[l:])
		if i < 0 {
			i = strings.IndexAny(s[l:], s[:l]) + l
		}
		if s[i] < 96 {
			// A - Z | 65 - 90
			sum += int(s[i]) - 38
		} else {
			// a - z | 97 - 122
			sum += int(s[i]) - 96
		}
	}
	fmt.Println(sum)
}

func part2(buf io.Reader) {
	sum := 0
	var l1, l2, l3 int
	var s1, s2, s3 string
	i := 0
	cross := make([]rune, 0, 16)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		s1 = scanner.Text()
		scanner.Scan()
		s2 = scanner.Text()
		scanner.Scan()
		s3 = scanner.Text()

		l1 = len(s1)
		l2 = len(s2)
		l3 = len(s3)

		if l1 > l2 && l1 > l3 {
			s1, s3 = s3, s1
			l1, l3 = l3, l1
		} else if l2 > l1 && l2 > l3 {
			s2, s3 = s3, s2
			l2, l3 = l3, l2
		}

		cross = cross[:0]

		for _, r := range s1 {
			if strings.IndexByte(s2, byte(r)) >= 0 {
				cross = append(cross, r)
			}
		}

		i = strings.IndexAny(s3, string(cross))

		if s3[i] < 96 {
			// A - Z | 65 - 90
			sum += int(s3[i]) - 38
		} else {
			// a - z | 97 - 122
			sum += int(s3[i]) - 96
		}
	}
	fmt.Println(sum)
}

func Variant1() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	part1(f)
	f, err = os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	part2(f)
}

func Variant2() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &buffer.Buffer{}
	io.Copy(buf, f)
	part1(buf)
	buf.Seek0()
	part2(buf)
}
