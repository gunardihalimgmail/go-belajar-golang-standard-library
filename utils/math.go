package utils

import (
	"fmt"
	"math"
)

func Math() {
	val1 := 157.4359
	temp1 := math.Floor(val1)
	temp2 := math.Ceil(val1)
	temp3 := math.Round(val1)
	temp4 := math.Max(10, 50)
	temp5 := math.Min(10, 50)

	fmt.Println("[Math Floor]", temp1)
	fmt.Println("[Math Ceil]", temp2)
	fmt.Println("[Math Round]", temp3)
	fmt.Println("[Math Max]", temp4)
	fmt.Println("[Math Min]", temp5)
}
