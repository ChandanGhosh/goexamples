package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := fmt.Sprintln("Enter number to find its factorial:")
	printer(prompt)
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			printer("Invalid entry! enter integer only", prompt)
		}
		printer(fmt.Sprintf("Factorial of %v is %v\n", num, factorial(num)), prompt)
	}

}

func printer(lines ...string) {
	for _, v := range lines {
		fmt.Println(v)
	}
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
