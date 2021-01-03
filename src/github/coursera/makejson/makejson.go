/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON
object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput(promptString string) string {
	fmt.Print(promptString)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	person := make(map[string]string)
	pName := getInput("Enter a name: ")
	person["name"] = pName
	pAddress := getInput("Enter an address: ")
	person["address"] = pAddress
	pJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(pJson))
}
