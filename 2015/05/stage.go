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

func isNice2(in string) bool {
	var inLen = len(in)
	var dblLetter bool = false
	var palindrome3 bool = false
	var pairs = make(map[[2]rune]int)
	var runes []rune = []rune(in)

	if inLen < 3 {
		return false
	}

	for i := 1; i < inLen; i++ {
		pair := [2]rune{runes[i-1], runes[i]}
		pos := pairs[pair]

		/* Find pairs of letters. */
		if !dblLetter && pos > 0 && pos <= i-2 {
			dblLetter = true
		} else if pos == 0 {
			/* Add pair in map with value the pos of the last letter. */
			pairs[pair] = i
		}

		if !palindrome3 && i >= 2 && runes[i-2] == runes[i] {
			palindrome3 = true
		}
	}

	if palindrome3 == true && dblLetter == true {
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
	var nbNice2 int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str string = scanner.Text()
		var nice bool
		var nice2 bool

		fmt.Printf("%s: ", str)
		if nice = isNice(str); nice {
			nbNice++
		}
		fmt.Printf("\t%t", nice)
		if nice2 = isNice2(str); nice2 {
			nbNice2++
		}
		fmt.Printf("\t%t\n", nice2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of nice strings: %d\n", nbNice)
	fmt.Printf("Number of nice2 strings: %d\n", nbNice2)
}
