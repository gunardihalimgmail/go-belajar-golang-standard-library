package main

import (
	"belajar-golang-standard-library/utils"
	"fmt"
)

func main() {
	// Package 'fmt'
	fmt.Println("Hello My Country")

	var firstName = "Jack"
	var lastName = "Strong"
	var age = 30
	fmt.Printf("Hello %s %s, %d years\n", firstName, lastName, age)

	fmt.Println("--------- Package Errors")
	utils.UtilsErrors()
	fmt.Println("--------- ------")

	fmt.Println("--------- Package OS")
	utils.Os()
	fmt.Println("--------- ------")

	fmt.Println("--------- Package Flag")
	utils.Flags()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Strings")
	utils.Strings()
	fmt.Println("--------- ------")

	fmt.Println("--------- Package StrConv")
	utils.StrConv()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Math")
	utils.Math()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Container List")
	utils.ContainerList()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Container Ring")
	utils.ContainerRing()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Sort")
	utils.SortImpl()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Time")
	utils.Time()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Reflect")
	utils.Reflect()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package RegExp")
	utils.Regexp()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Encoding")
	utils.Encoding()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package CSV Reader")
	utils.CsvReader()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package CSV Writer")
	utils.CsvWriter()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Slices")
	utils.Slices()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Path")
	utils.Path()
	utils.FilePath()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Bufio Reader")
	utils.BufioRead()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package Bufio Writer")
	utils.BufioWriter()
	fmt.Println("--------- ------")

	fmt.Println("\n--------- Package File Manipulation (OS)")
	utils.OsCreateFile()
	fmt.Println("--------- ------")

}
