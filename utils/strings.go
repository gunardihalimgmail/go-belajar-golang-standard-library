package utils

import (
	"fmt"
	"strings"
)

func Strings() {
	// temp := "Golang Programming"
	temp1 := strings.Contains("Golang Programming", "lang")
	temp2 := strings.Trim("    Golang Programming ", " ")
	temp3 := strings.ToLower("Golang Programming ")
	temp4 := strings.ToUpper("Golang Programming ")
	temp5 := strings.ToTitle("golang programming ")
	temp6 := strings.ReplaceAll("Golang Programming Golang", "Golang", "Javascript")
	temp7 := strings.Split("Golang_Programming", "_")

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)
	fmt.Println(temp4)
	fmt.Println(temp5)
	fmt.Println(temp6)
	fmt.Println(temp7)

}
