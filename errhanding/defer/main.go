package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writerFile("fib.txt")
}

func writerFile(filename string)  {
	//file, err := os.Create(filename)//创建一个文件,建文件的时候就要想着关掉文件

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)//创建bufio缓冲writer的时候要想着flush
	defer writer.Flush()
	for i := 1; i<=20; i++ {
		fmt.Fprintln(writer, fibonacci(i))
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	return
	fmt.Println(3)
}

func fibonacci(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}
	
	return fibonacci(n-2) + fibonacci(n-1)
}
