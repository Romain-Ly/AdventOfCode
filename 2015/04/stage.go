package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
)

func md5Build(input string) []byte {
	var h = md5.New()

	io.WriteString(h, input)
	md5bin := h.Sum(nil)

	var dst = make([]byte, hex.EncodedLen(len(md5bin)))
	hex.Encode(dst, md5bin)

	return dst
}

func startsWith5Zeros(input []byte) bool {
	if len(input) <= 5 {
		return false
	}

	for i := 0; i < 5; i++ {
		if input[i] != '0' {
			return false
		}
	}

	return true
}

func main() {
	var secret string = os.Args[1]
	var find bool = false

	for i := 1; find == false; i++ {
		output := md5Build(secret + strconv.Itoa(i))
		if i%10000 == 0 {
			fmt.Println(i)
		}
		find = startsWith5Zeros(output)
		if find {
			fmt.Printf("%d: %s\n", i, output)
		}
	}
}
