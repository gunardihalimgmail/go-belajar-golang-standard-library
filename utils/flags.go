package utils

import (
	"flag"
	"fmt"
)

func Flags() {

	// ---Cara Penggunaan
	// go run .\fmt.go -username="Gunardi Halim" -password=Rahasia -host="123.123.123.123" -port=2341

	var username *string = flag.String("username", "root", "Input Username")
	var password *string = flag.String("password", "rahasia", "Input Password")
	var host *string = flag.String("host", "123.84.91.10", "Input Host")
	var port *int = flag.Int("port", 7789, "Input Port")

	flag.Parse()

	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Host :", *host)
	fmt.Println("Port :", *port)
}
