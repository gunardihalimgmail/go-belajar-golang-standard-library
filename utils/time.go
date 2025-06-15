package utils

import (
	"fmt"
	"time"
)

func Time() {
	var now time.Time = time.Now()
	fmt.Println("Now :", now)
	fmt.Println("Now (Local) :", now.Local())
	fmt.Println("Now (Add) :", now.AddDate(10, 1, 1).Add(10000000000))    // Add (10 detik) : konversi sampai nanoseconds. 1 detik = 1000 milisec, 1 milsec = 1000 micro, 1 micro = 1000 nonasec
	fmt.Println("Now (Add) :", now.AddDate(10, 1, 1).Add(10*time.Second)) // Add Cara 2 (10 detik)

	var date time.Time = time.Date(2025, 7, 9, 10, 10, 10, 0, time.UTC)
	fmt.Println("Date :", date)

	// Konversi tanggal string ke date sesuai format
	formatter := "2006-01-02 15:04:05" // harus tanggal yang ini 2 Jan 2006
	value := "2025-05-03 01:20:30"
	// value := "2025-05-03T01:20:30Z"	// time.RFC3339

	// parse, err := time.Parse(time.RFC3339, value) //cara 1
	parse, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("[Time Parse]", parse)
	}

	// Hitung Durasi Waktu
	var time1 time.Duration = 100 * time.Second // 10 detik
	var time2 time.Duration = 10 * time.Millisecond
	durasi := time1 - time2
	fmt.Println("[Time Parse]", durasi)
}
