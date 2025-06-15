package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func BufioRead() {
	input := strings.NewReader("This is new string\nThis is New Line")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
