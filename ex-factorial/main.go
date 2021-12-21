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
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
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

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func factorial_improved(n int) int64 {
	if n <= 1 {
		return 1
	}
	var fact int64 = 1
	for i := 1; i <= n; i++ {
		fact *= int64(i)
	}
	return fact
}
