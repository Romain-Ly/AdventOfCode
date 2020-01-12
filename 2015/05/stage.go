package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isNice(in string) bool {
	var inLen int = len(in)
	var nbVowels int = 0
	var previous byte = 0
	var dbleLetter bool = false

	for i := 0; i < inLen; i++ {
		switch in[i] {
		case 'a', 'e', 'i', 'o', 'u':
			nbVowels++

		/* Bad strings */
		case 'b':
			if i >= 1 && previous == 'a' {
				return false
			}
		case 'd':
			if i >= 1 && previous == 'c' {
				return false
			}

		case 'q':
			if i >= 1 && previous == 'p' {
				return false
			}

		case 'y':
			if i >= 1 && previous == 'x' {
				return false
			}
		}

		/* Double Letters */
		if i >= 1 && dbleLetter == false && previous == in[i] {
			dbleLetter = true
		}
		previous = in[i]

	}

	if nbVowels >= 3 && dbleLetter == true {
		return true
	}
	return false
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nbNice int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str string = scanner.Text()
		var nice bool

		fmt.Printf("%s: ", str)
		if nice = isNice(str); nice {
			nbNice++
		}
		fmt.Printf("%t\n", nice)
	}

	fmt.Printf("Number of nice strings: %d\n", nbNice)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
