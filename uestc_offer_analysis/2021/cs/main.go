package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "utc_2021_cs_scores.txt"
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	steps := map[int]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		step := score - score%10

		if step > 400 {
			step = 400
		}
		steps[step]++
	}

	for step := 300; step <= 400; step += 10 {
		fmt.Printf("%d,%d\n", step, steps[step])
	}
}
