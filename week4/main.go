package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func encode(input string) string {
	var encoded strings.Builder
	var row, col int
	encoded.Grow(len(input) * 3)

	for _, char := range strings.ToUpper(input) {
		if char >= 'A' && char <= 'Z' {
			offset := int(char - 'A')
			row, col = offset/6, offset%6

		} else if char >= '0' && char <= '9' {
			offset := int(char-'0') + 24
			row, col = offset/6, 2+offset%6
		}
		encoded.WriteByte(byte(row + '1'))
		encoded.WriteByte(byte(col + '1'))
	}

	return encoded.String()
}

func decode(input string) string {
	var decoded strings.Builder
	decoded.Grow(len(input) / 2)

	for i := 0; i < len(input); i += 2 {
		pair := input[i : i+2]
		index, _ := strconv.Atoi(pair)

		if index < 53 {
			decoded.WriteByte('A' + byte((index/10-1)*6+index%10-1))
		} else {
			decoded.WriteByte('0' + byte((index/10-6+1)*6+index%10-3))
		}
	}

	return decoded.String()
}

func main() {
	input := flag.String("t", "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "text to cipher")
	flag.Parse()

	encoded := encode(*input)
	fmt.Printf("Encoded: %s\n", encoded)

	decoded := decode(encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
