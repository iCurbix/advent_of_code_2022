package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cumSum := 0
	maxSum := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			cumSum += num
		} else {
			if cumSum > maxSum {
				maxSum = cumSum
			}
			cumSum = 0
		}
	}
	if cumSum > maxSum {
		maxSum = cumSum
	}
	fmt.Println(maxSum)
}

func part2() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	max3 := []int{0, 0, 0}
	cumSum := 0

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err == nil {
			cumSum += num
		} else {
			for i, v := range max3 {
				if cumSum > v {
					cumSum, max3[i] = max3[i], cumSum
				}
			}
			cumSum = 0
		}
	}
	fmt.Println(max3[0] + max3[1] + max3[2])
}

func main() {
	part1()
	part2()
}
