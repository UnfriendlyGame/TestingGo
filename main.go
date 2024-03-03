package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	numOfArgs := len(os.Args)
	if numOfArgs != 1 && numOfArgs != 3 {
		fmt.Println("Wrong Number Of ARGUEMNTS")
		os.Exit(-1)
	}
	if numOfArgs == 1 {
		fmt.Println("enter arguments")
	}
	if numOfArgs == 3 {
		result := strconv.Atoi(os.Args[1]) + strconv.Atoi(os.Args[2])
		fmt.Println("%v + %v = %v", os.Args[1], os.Args[2], result)
	}
}
