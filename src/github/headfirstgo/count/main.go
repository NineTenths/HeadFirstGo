// count tallies the nunmber of times each line occurs within a file
package main

import (
	"fmt"
	"github/headfirstgo/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrigs("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
