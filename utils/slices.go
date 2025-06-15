package utils

import (
	"fmt"
	"slices"
)

func Slices() {
	names := []string{"Jodi", "Misha", "Agung", "Rahmat"}
	values := []int{100, 90, 85, 95}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Misha"))
	fmt.Println(slices.Contains(values, 85))
	fmt.Println(slices.Index(names, "Agung"))
	fmt.Println(slices.Index(values, 95))
}
