package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func validateAge(s string) (int, error) {
	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	if n < 0 {
		return 0, fmt.Errorf("negative")
	}

	return n, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, err := validateAge(sc.Text())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", age)
	}
}
