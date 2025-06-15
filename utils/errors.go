package utils

import (
	"errors"
	"fmt"
)

var (
	ValidationError  = errors.New("Validation Error")
	NoDataFoundError = errors.New("No Data Found Error")
)

func saveData(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "jack" {
		return NoDataFoundError
	}
	return nil
}

func UtilsErrors() {
	err := saveData("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println(err.Error())
		} else if errors.Is(err, NoDataFoundError) {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Sukses")
	}

}
