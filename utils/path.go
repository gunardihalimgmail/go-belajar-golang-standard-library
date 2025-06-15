package utils

import (
	"fmt"
	"path"
	"path/filepath"
)

// Beda Path dan FilePath yaitu kalau filepath bisa mendeteksi OS yang digunakan
// kalau OS Unix / Mac pakai "/", sedangkan Windows pakai "\"
func Path() {
	fmt.Println("--> Path <--")
	fmt.Println(path.Dir("hello/indo.go"))              // hello
	fmt.Println(path.Base("hello/indo.go"))             // indo.go
	fmt.Println(path.Ext("hello/indo.go"))              // .go
	fmt.Println(path.Join("hello", "Great", "Indo.go")) // hello/Great/Indo.go
}

func FilePath() {
	// FilePath dapat menyesuaikan format OS yang dipakai, kalau Windows outputnya '\'
	fmt.Println("\n--> FilePath <--")
	fmt.Println(filepath.Dir(`hello\indo.go`))              // nama direktori : hello
	fmt.Println(filepath.Base("hello/indo.go"))             // nama file : indo.go
	fmt.Println(filepath.Ext("hello/indo.go"))              // .go
	fmt.Println(filepath.IsAbs("hello/indo.go"))            // Abs : false karena tidak dimulai full dari 'C:\hello\...'
	fmt.Println(filepath.IsLocal("hello/indo.go"))          // Local : true karena file berada di pada lokal
	fmt.Println(filepath.Join("hello", "Great", "Indo.go")) // Join : 'hello\Great\Indo.go'
}
