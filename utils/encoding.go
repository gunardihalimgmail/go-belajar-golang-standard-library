package utils

import (
	"encoding/base64"
	"fmt"
)

func Encoding() {
	temp := "Gunardi Halim"
	encoding := base64.StdEncoding.EncodeToString([]byte(temp)) // encode ke base64
	fmt.Println("Encode :", encoding)

	decoding, err := base64.StdEncoding.DecodeString(encoding)
	if err != nil {
		fmt.Println("Error Decoding :", err)
	} else {
		fmt.Println("Decode :", string(decoding))
	}
}
