package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(f *os.File, b1 []byte) (n int) {
	n1, _ := f.Read(b1)
	return n1
}

func main() {
	fd, err := os.Open(os.Args[1])
	check(err)

	b1 := make([]byte, 1)
	var res int = 0
	var pos int = 0

	for read(fd, b1) > 0 {
		if b1[0] == 41 {
			res -= 1
		} else if b1[0] == 40 {
			res += 1
		}

		pos++
		if res == -1 {
			fmt.Printf("basement at pos %d\n", pos)
		}
	}
	fmt.Printf("%d\n", res)

}
