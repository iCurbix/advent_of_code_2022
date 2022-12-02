package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
A - 65
B - 66
C - 67                       A  B  C
X - 88      23  22  21    X  D  L  W    Z  X  Y
Y - 89      24  23  22    Y  W  D  L    X  Y  Z
Z - 90      25  24  23    Z  L  W  D    Y  Z  X
*/

func score(a, b uint8) uint8 {
	x := b - a
	var r uint8
	if x == 23 {
		r = 3
	} else if x == 24 || x == 21 {
		r = 6
	} else {
		r = 0
	}
	return r + b - 87
}

func getMove(a, b uint8) uint8 {
	return (3-(((a+2)%3+(a-65))%3)+b-88)%3 + 88
}

func part1() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var a, b uint8
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a = scanner.Text()[0]
		b = scanner.Text()[2]
		sum += int(score(a, b))
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var a, b, c uint8
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a = scanner.Text()[0]
		b = scanner.Text()[2]
		c = getMove(a, b)
		sum += int(score(a, c))
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
