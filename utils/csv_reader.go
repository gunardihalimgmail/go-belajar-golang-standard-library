package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func CsvReader() {
	data := "Jakarta, DKI, Indonesia\n" +
		"Beijing, Metro, China\n" +
		"Hamburg, Great, Germany"

	datacsv := csv.NewReader(strings.NewReader(data))

	for {
		record, err := datacsv.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

}
