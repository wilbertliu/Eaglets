package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	// If there's error, print the error and stop the execution.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Otherwise print the user input.
	fmt.Println(text)
}
