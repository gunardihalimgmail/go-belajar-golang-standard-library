// Struktur Data Double Linked List
package utils

import (
	"container/list"
	"fmt"
)

func ContainerList() {
	var student *list.List = list.New()
	student.PushBack("Gunardi")
	student.PushBack("Halim")
	student.PushBack("Javier")

	dataDepan := student.Front()
	fmt.Println(dataDepan.Value)

	dataNext := dataDepan.Next()
	fmt.Println(dataNext.Value)

	dataNext2 := dataNext.Next()
	fmt.Println(dataNext2.Value)

	fmt.Println("")

	for e := student.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
