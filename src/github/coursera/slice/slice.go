/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user
to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the
slice, and prints the contents of the slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character ‘X’ instead of an integer.

1/1/2021
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// The following line is how to create a slice with a length of 3, but this makes no sense in the context of
	// this program since the length can and probably will be longer than 3.
	//numbers:= make([]int,3)
	// The following line is how I would create the slice in an actual program
	var numbers []int
	fmt.Print("Enter an integer (or X to quit): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	for input != "X" {
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
		sort.Ints(numbers)
		fmt.Println(numbers)
		fmt.Print("Enter an integer (or X to quit): ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
	}
}
