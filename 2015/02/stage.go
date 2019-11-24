package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var prism []string = strings.Split(scanner.Text(), "x")
		var sq [3]int
		for i, str := range prism {
			sq[i], _ = strconv.Atoi(str)
		}

		var a int = sq[0] * sq[1]
		var b int = sq[1] * sq[2]
		var c int = sq[0] * sq[2]

		var slack int = a
		if b < slack {
			slack = b
		}
		if c < slack {
			slack = c
		}

		total += 2*(a+b+c) + slack
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
