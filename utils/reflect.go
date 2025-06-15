// Fungsi untuk mengetahui meta data dari suatu struct
package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name    string `required:"true" max:"10"` // Struct Tag ditandai Back Tick (`)
	Address string `required:"true" max:"20"`
	City    string `required:"true" max:"30"`
}

func isValid(data any) bool {

	result := true
	// Cek apakah data kosong jika ada required
	var dataType reflect.Type = reflect.TypeOf(data)
	for i := 0; i < dataType.NumField(); i++ {
		var dataField reflect.StructField = dataType.Field(i)

		dataTag := dataField.Tag.Get("required")

		if dataTag != "" {

			res, err := strconv.ParseBool(dataTag)
			if err != nil {
				fmt.Println("Bukan Boolean (True / False) :", err)
				result = false
			} else {
				if res { // jika required true
					dataValue := reflect.ValueOf(data).Field(i).Interface() // ambil data value
					if dataValue == "" {
						result = false
					}
				}
			}

		}
	}
	return result
}

func Reflect() {
	person := Person{Name: "Pirates", Address: "Bougenville Street", City: "Jakarta"}
	var reflectType reflect.Type = reflect.TypeOf(person)

	for i := 0; i < reflectType.NumField(); i++ {

		var field reflect.StructField = reflectType.Field(i) // meta data field

		value := reflect.ValueOf(person).Field(i).Interface() // isi value datanya

		fmt.Println("'", field.Name, field.Index, "'", "dengan Type ='", field.Type, "' Value -> '", value, "'")

	}

	// Validasi Reflect
	fmt.Println("\n------ Validation Reflect ------")
	person2 := Person{Name: "Pirates", Address: "Merdeka", City: "Jakarta"}
	fmt.Println("Valid : ", isValid(person2))

}
