package main

import (
	"fmt"
	"time"
)

func main() {
	newYearUnix := time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Printf("\r%v s", newYearUnix-time.Now().Unix())
	}
}
