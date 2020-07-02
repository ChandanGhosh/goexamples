// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)

	scnanner := bufio.NewScanner(os.Stdin)

	fmt.Println("I am listening...")
	for scnanner.Scan() {

		if scnanner.Text() == "" {
			break
		}
		count[scnanner.Text()]++

	}

	for line, cnt := range count {
		fmt.Printf("%s:\t%d\n", line, cnt)
	}
}
