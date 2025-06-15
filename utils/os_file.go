package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createFile() error {
	// * 066 adalah 'permission' dari chmod unix, otorisasi nya bisa read write untuk "owner, group, public"
	// akses situs "chmod-calculator" untuk mengetahui kode permission unix ini
	file, err := os.OpenFile("sample.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString("This is Sample Log")
	return nil
}

func addToFile(nama_file, message string) error {
	file, err := os.OpenFile(nama_file, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile() (string, error) {
	file, err := os.OpenFile("sample.log", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Errorrr :", err.Error())
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, errRead := reader.ReadLine()
		if errRead == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func OsCreateFile() {
	fmt.Println("\n------ Create File")
	createFile()

	fmt.Println("\n------ Add File")
	addToFile("sample.log", "\nThis is Log 1")
	addToFile("sample.log", "\nThis is Log 2")

	fmt.Println("\n------ Read File")
	rec, _ := readFile()
	fmt.Println(rec)
}
