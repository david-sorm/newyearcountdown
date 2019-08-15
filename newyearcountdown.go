package main

import (
	"fmt"
	"time"
)

func main() {
	newYearUnix := time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	for {
		time.Sleep(1 * time.Second)
		currentUnix := time.Now().Unix()
		fmt.Printf("\r%v s", newYearUnix-currentUnix)
	}
}
