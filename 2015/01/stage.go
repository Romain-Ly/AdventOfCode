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
	f, err := os.Open(os.Args[1])
	check(err)
	b1 := make([]byte, 1)

	var res int = 0
	for read(f, b1) > 0 {
		if b1[0] == 41 {
			res -= 1
		} else if b1[0] == 40 {
			res += 1
		}
	}
	fmt.Printf("%d\n", res)

}
