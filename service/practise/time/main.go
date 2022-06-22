package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	var s2 = time.Tick(2 * time.Second)
	var s3 = time.Tick(7 * time.Second)
	for {
		select {
		case <-s2:
			fmt.Println("2 second over:", time.Now().Second())
		case <-s3:
			fmt.Println("5 second over, timeover", time.Now().Second())
			return
		}
	}
}
