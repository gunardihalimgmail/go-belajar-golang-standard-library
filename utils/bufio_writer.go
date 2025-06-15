package utils

import (
	"bufio"
	"os"
)

func BufioWriter() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Here is a Math Lesson\n")
	_, _ = writer.WriteString("Here is a Good Subject\n")
	writer.Flush()
}
