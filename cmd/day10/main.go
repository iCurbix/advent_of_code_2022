package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	reg  int
	halt bool
}

type instruction string

const (
	noop instruction = "noop"
	addx instruction = "addx"
)

func (c *cpu) do(instr instruction, args ...string) bool {
	switch instr {
	case noop:
		return true
	case addx:
		if !c.halt {
			c.halt = true
			return false
		}
		q, _ := strconv.Atoi(args[0])
		c.reg += q
		c.halt = false
		return true
	default:
		return true
	}
}

func main() {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var s1, s2 string
	cpu := &cpu{reg: 1}
	sum := 0
	i := 1
	pos := 0
	fin := false
	for scanner.Scan() {
		s1, s2, _ = strings.Cut(scanner.Text(), " ")
		for {
			pos = (i - 1) % 40
			if pos > cpu.reg-2 && pos < cpu.reg+2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if pos == 39 {
				fmt.Println()
			}
			fin = cpu.do(instruction(s1), s2)
			i++
			if (i-20)%40 == 0 {
				sum += i * cpu.reg
			}
			if fin {
				break
			}
		}
	}
	fmt.Println(sum)
}
