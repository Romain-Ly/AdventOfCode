package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeRibbon(a int, b int, c int) int {
	var x int = a
	var y int = b
	var maxSide int = c

	if x > maxSide {
		maxSide, x = x, maxSide /* swap */
	}
	if y > maxSide {
		maxSide, y = y, maxSide
	}

	return 2*(x+y) + a*b*c
}

func computeWrappingPaper(a int, b int, c int) int {
	var x int = a * b
	var y int = b * c
	var z int = a * c

	var slack int = x
	if y < slack {
		slack = y
	}
	if z < slack {
		slack = z
	}

	return 2*(x+y+z) + slack
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wrapping int = 0
	var ribbon int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var prism []string = strings.Split(scanner.Text(), "x")
		var sq [3]int
		for i, str := range prism {
			sq[i], _ = strconv.Atoi(str)
		}
		wrapping += computeWrappingPaper(sq[0], sq[1], sq[2])
		ribbon += computeRibbon(sq[0], sq[1], sq[2])
	}
	fmt.Printf("wrapping paper: %d\n", wrapping)
	fmt.Printf("rbbon length: %d\n", ribbon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
