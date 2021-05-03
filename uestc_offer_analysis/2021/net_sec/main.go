package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "scores.txt"
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	n := 0
	max := 0
	min := 0
	steps := map[int]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum += score
		n++

		if max == 0 || score > max {
			max = score
		}
		if min == 0 || score < min {
			min = score
		}

		step := score - score%10
		if step > 400 {
			step = 400
		}
		steps[step]++
	}

	fmt.Printf("count:%d\n", n)
	fmt.Printf("avarage:%d\n", sum/n)
	fmt.Printf("min:%d\n", min)
	fmt.Printf("max:%d\n", max)
	fmt.Println()

	for step := 300; step <= 400; step += 10 {
		fmt.Printf("%d\n", steps[step])
	}
}
