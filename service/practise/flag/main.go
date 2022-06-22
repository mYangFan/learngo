package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("e", "", "the environment")

	flag.Parse()

	fmt.Println(*s)
}
