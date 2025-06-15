package utils

import (
	"encoding/csv"
	"os"
)

func CsvWriter() {
	data := csv.NewWriter(os.Stdout) // output ke Terminal
	_ = data.Write([]string{"China"})
	_ = data.Write([]string{"Japan"})
	_ = data.Write([]string{"Korea"})
	data.Flush()
}
