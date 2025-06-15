// Konsep Ring yang berputar tanpa henti. jika sampai ke angka terakhir, maka akan kembali ke angka awal
// Contoh  Data 1 -> 2 -> 3 -> 4 -> 5 -> 1 -> 2 -> 3 ->...

package utils

import (
	"container/ring"
	"fmt"
)

func ContainerRing() {
	var data *ring.Ring = ring.New(5)

	// ** Manual

	data.Value = "1"
	// fmt.Println(data.Value)

	data = data.Next()
	data.Value = "2"
	// fmt.Println(data.Value)

	data = data.Next()
	data.Value = "3"
	// fmt.Println(data.Value)

	data = data.Next()
	data.Value = "4"

	data = data.Next()
	data.Value = "5"

	data.Do(func(v any) {
		fmt.Println(v)
	})

}
