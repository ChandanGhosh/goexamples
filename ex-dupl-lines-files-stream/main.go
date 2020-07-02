package main

import (
	"bufio"
	"fmt"
	"os"
)

var count map[string]int

func main() {
	count = make(map[string]int)
	files := os.Args[1:]

	if len(files) != 0 {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("error in opening file, %s", err.Error())
				continue
			}

			countLines(f, count)
			f.Close()
		}

	} else {
		countLines(os.Stdin, count)
	}

	for line, cnt := range count {
		fmt.Printf("%s\t:%d\n", line, cnt)
	}

}

func countLines(f *os.File, count map[string]int) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		count[scanner.Text()]++
	}

}
