package utils

import (
	"fmt"
	"strconv"
)

func StrConv() {

	var tempStrInt = "12734895"
	tempParseInt, errInt := strconv.ParseInt(tempStrInt, 10, 64)
	if errInt != nil {
		fmt.Println("Error :", errInt.Error())
	} else {
		fmt.Println("[Parse Int]", tempParseInt)
	}
	var tempStrBool = "true"
	tempParseBool, errBool := strconv.ParseBool(tempStrBool)
	if errInt != nil {
		fmt.Println("Error :", errBool.Error())
	} else {
		fmt.Println("[Parse Bool]", tempParseBool)
	}

	var tempInt int64 = 5
	tempFormatInt := strconv.FormatInt(tempInt, 10) // Konversi int64 ke decimal ke string

	tempBool := true
	tempFormatBool := strconv.FormatBool(tempBool) // Konversi bool ke string

	tempFloat := 89233.1425034458648691
	// -1 : default sesuai panjang decimal
	// 'f' : no exponential
	tempFormatFloat := strconv.FormatFloat(tempFloat, 'f', -1, 64) // Konversi bool ke string

	fmt.Println("[Format Int]", tempFormatInt)
	fmt.Println("[Format Bool]", tempFormatBool)
	fmt.Println("[Format Float]", tempFormatFloat)
}
