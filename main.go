package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Print(">>> ")
		}

		program, err := Compile(input)
		if err != nil {
			fmt.Printf("error while compiling: %s", err.Error())
		}
		output, err := Execute(program)
		if err != nil {
			fmt.Printf("error while executing: %s", err.Error())
		}
		fmt.Println(output)

		fmt.Print(">>> ")
	}
}
