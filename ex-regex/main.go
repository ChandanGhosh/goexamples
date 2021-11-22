package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 12 shares of MSFT for $234.54
// 10 shares of TSLA for $692.3

var matchRegEx = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.d+)?)`)

type Transaction struct {
	Symbole string
	Volume  int
	Price   float64
}

func parseLine(line string) (Transaction, error) {
	matches := matchRegEx.FindStringSubmatch(line)
	if matches == nil {
		return Transaction{}, fmt.Errorf("bad line: %q", line)
	}
	var t Transaction
	t.Symbole = matches[2]
	t.Volume, _ = strconv.Atoi(matches[1])
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}

func main() {
	line := "12 shares of MSFT for $234.54"
	t, err := parseLine(line)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// when printing the struct, +v adds the field names
	fmt.Printf("%+v\n", t)
}