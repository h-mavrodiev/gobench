package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Take input as path for configuration
	// 	Endpoints
	// 	Payload
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter config  path: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	// Take input for number of requests per endpoint

}
