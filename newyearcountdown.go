package main

import (
	"fmt"
	"time"
)

func refresh(newYearUnix int64) {
	currentUnix := time.Now().Unix()
	go fmt.Printf("\r%v s", newYearUnix-currentUnix)
}

func main() {
	newYearUnix := time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	for {
		time.Sleep(1 * time.Second)
		go refresh(newYearUnix)
	}
}
