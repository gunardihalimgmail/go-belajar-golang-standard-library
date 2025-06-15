package utils

import (
	"fmt"
	"os"
)

func Os() {
	args := os.Args
	for i, arg := range args {
		fmt.Println(i, " -> ", arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error :", err)
	}
}
