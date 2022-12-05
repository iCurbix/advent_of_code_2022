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

func appendStacks(stacks [][]uint8, num int, s string) {
	// 4n+1
	for i := 0; i < num; i++ {
		if s[4*i+1] != ' ' {
			stacks[i] = append(stacks[i], s[4*i+1])
		}
	}
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	stacksNum := (len(s) + 1) / 4
	stacks := make([][]uint8, stacksNum)
	appendStacks(stacks, stacksNum, s)
	for scanner.Scan() {
		s = scanner.Text()
		if s[1] == '1' {
			break
		}
		appendStacks(stacks, stacksNum, s)
	}

	for _, v := range stacks {
		for i := 0; i < len(v)/2; i++ {
			v[i], v[len(v)-i-1] = v[len(v)-i-1], v[i]
		}
	}

	scanner.Scan()

	var num, from, to int
	var splitted []string
	for scanner.Scan() {
		splitted = strings.Split(scanner.Text(), " ")
		num, _ = strconv.Atoi(splitted[1])
		from, _ = strconv.Atoi(splitted[3])
		to, _ = strconv.Atoi(splitted[5])

		for i := len(stacks[from-1]) - 1; i >= len(stacks[from-1])-num; i-- {
			stacks[to-1] = append(stacks[to-1], stacks[from-1][i])
		}
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-num]
	}

	for _, v := range stacks {
		fmt.Printf("%c", v[len(v)-1])
	}
}

func part2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	stacksNum := (len(s) + 1) / 4
	stacks := make([][]uint8, stacksNum)
	appendStacks(stacks, stacksNum, s)
	for scanner.Scan() {
		s = scanner.Text()
		if s[1] == '1' {
			break
		}
		appendStacks(stacks, stacksNum, s)
	}

	for _, v := range stacks {
		for i := 0; i < len(v)/2; i++ {
			v[i], v[len(v)-i-1] = v[len(v)-i-1], v[i]
		}
	}

	scanner.Scan()

	var num, from, to int
	var splitted []string
	for scanner.Scan() {
		splitted = strings.Split(scanner.Text(), " ")
		num, _ = strconv.Atoi(splitted[1])
		from, _ = strconv.Atoi(splitted[3])
		to, _ = strconv.Atoi(splitted[5])

		stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-num:]...)
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-num]
	}

	for _, v := range stacks {
		fmt.Printf("%c", v[len(v)-1])
	}
}

func main() {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	fmt.Println()
	part2(buf)
}
