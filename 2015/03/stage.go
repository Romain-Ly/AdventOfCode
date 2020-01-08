package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getStrPos(xpos int, ypos int) (coord string) {
	return strconv.Itoa(xpos) + "." + strconv.Itoa(ypos)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int = 0
	var xpos int = 0
	var ypos int = 0
	var maps = make(map[string]bool)
	maps[getStrPos(0, 0)] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, arrow := range scanner.Text() {
			//fmt.Printf("%c", arrow)
			switch arrow {
			case 60: // <
				xpos--
			case 62: // >
				xpos++
			case 94: // ^
				ypos++
			case 118: // v
				ypos--
			}

			/* Check if coordinates has been visited. */
			var coord string = getStrPos(xpos, ypos)
			visited := maps[coord]
			if !visited {
				maps[coord] = true
				total++
			}
		}
	}
	fmt.Println(len(maps))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
