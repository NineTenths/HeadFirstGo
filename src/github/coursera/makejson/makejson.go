package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string)
	fmt.Print("Enter a name: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	person["name"] = input
	fmt.Println("You entered: " + person["name"])
}
