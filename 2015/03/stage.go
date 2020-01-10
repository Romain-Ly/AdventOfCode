package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type robot struct {
	name string
	xpos int
	ypos int
}

func getStrPos(xpos int, ypos int) (coord string) {
	return strconv.Itoa(xpos) + "." + strconv.Itoa(ypos)
}

func NewRobot(name string) robot {
	p := robot{name: name}
	p.xpos = 0
	p.ypos = 0
	return p
}

func MoveRobot(arrow rune, robot *robot) (pos string) {
	fmt.Printf("robot %s get %c\n", robot.name, arrow)
	switch arrow {
	case 60: // <
		robot.xpos--
	case 62: // >
		robot.xpos++
	case 94: // ^
		robot.ypos++
	case 118: // v
		robot.ypos--
	}

	/* Check if coordinates has been visited. */
	return getStrPos(robot.xpos, robot.ypos)
}

func VisitLoc(city map[string]bool, coord string) {
	visited := city[coord]
	if !visited {
		city[coord] = true
	}
}

/* ./stage <file> <nbOfRobots> */
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nbRobot, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	var robots = []robot{}
	for i := 0; i < nbRobot; i++ {
		n := NewRobot(strconv.Itoa(i))
		robots = append(robots, n)
	}

	var city = make(map[string]bool)
	VisitLoc(city, getStrPos(0, 0))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, arrow := range scanner.Text() {
			coord := MoveRobot(arrow, &robots[i%nbRobot])
			VisitLoc(city, coord)
		}
	}

	fmt.Printf("total of houses : %d\n", len(city))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
