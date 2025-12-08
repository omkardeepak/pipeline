package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting for test script...will take 3 seconds")
	time.Sleep(3 * time.Second)
	fmt.Println("Test script executed")

}
