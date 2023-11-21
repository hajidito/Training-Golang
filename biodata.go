package main

import (
	"assignment/method"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run biodata.go [nama teman]")
		return
	}
	namaTeman := os.Args[1]
	go method.TampilkanDataTeman(namaTeman)
	time.Sleep(time.Millisecond * 1)
}