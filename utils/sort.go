package utils

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortImpl() {

	data := []User{
		{"John", 30},
		{"Kety", 25},
		{"Zainal", 28},
		{"Yusuf", 36},
	}
	sort.Sort(UserSlice(data))

	fmt.Println(data)

}
