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

type operation uint8

const (
	mul operation = '*'
	add operation = '+'
)

var smallener = 1

type monke struct {
	items       []int
	testDiv     int
	a           int
	b           int
	op          operation
	m1          int
	m2          int
	inspections int
	skipDiv     bool
}

func (m *monke) Do(monkeys []*monke) {
	for _, it := range m.items {
		var a, b int
		if m.a == -1 {
			a = it
		} else {
			a = m.a
		}
		if m.b == -1 {
			b = it
		} else {
			b = m.b
		}
		switch m.op {
		case add:
			it = a + b
		case mul:
			it = a * b
		}
		m.inspections++
		if !m.skipDiv {
			it /= 3
		}
		if it%m.testDiv == 0 {
			monkeys[m.m1].items = append(monkeys[m.m1].items, it%smallener)
		} else {
			monkeys[m.m2].items = append(monkeys[m.m2].items, it%smallener)
		}
	}
	m.items = nil
}

func getMonkeys(r io.Reader) []*monke {
	scanner := bufio.NewScanner(r)
	var s string
	var it int
	var monkeys []*monke
	for scanner.Scan() {
		m := monke{}
		scanner.Scan()
		_, s, _ = strings.Cut(scanner.Text(), ": ")
		for _, v := range strings.Split(s, ", ") {
			it, _ = strconv.Atoi(v)
			m.items = append(m.items, it)
		}
		scanner.Scan()
		_, s, _ = strings.Cut(scanner.Text(), "= ")
		v := strings.Split(s, " ")
		vv, err := strconv.Atoi(v[0])
		if err != nil {
			m.a = -1
		} else {
			m.a = vv
		}
		m.op = operation(v[1][0])
		vv, err = strconv.Atoi(v[2])
		if err != nil {
			m.b = -1
		} else {
			m.b = vv
		}
		scanner.Scan()
		v = strings.Split(scanner.Text(), " ")
		vv, _ = strconv.Atoi(v[len(v)-1])
		m.testDiv = vv
		scanner.Scan()
		v = strings.Split(scanner.Text(), " ")
		vv, _ = strconv.Atoi(v[len(v)-1])
		m.m1 = vv
		scanner.Scan()
		v = strings.Split(scanner.Text(), " ")
		vv, _ = strconv.Atoi(v[len(v)-1])
		m.m2 = vv
		scanner.Scan()
		monkeys = append(monkeys, &m)
	}
	for _, m := range monkeys {
		smallener *= m.testDiv
	}
	return monkeys
}

func monkeyBusiness(monkeys []*monke) int {
	var max1, max2 int
	for _, m := range monkeys {
		if m.inspections > max2 {
			max2 = m.inspections
		}
		if max2 > max1 {
			max1, max2 = max2, max1
		}
	}
	return max1 * max2
}

func part1(r io.Reader) {
	monkeys := getMonkeys(r)
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.Do(monkeys)
		}
	}
	fmt.Println(monkeyBusiness(monkeys))
}

func part2(r io.Reader) {
	monkeys := getMonkeys(r)
	for _, m := range monkeys {
		m.skipDiv = true
	}
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.Do(monkeys)
		}
	}
	fmt.Println(monkeyBusiness(monkeys))
}

func main() {
	f, err := os.Open("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	smallener = 1
	part2(buf)
}
