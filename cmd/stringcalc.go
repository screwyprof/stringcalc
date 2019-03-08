package main

import (
	"fmt"
	"os"

	"github.com/screwyprof/stringcalc"
)

func main() {
	calc := stringcalc.StringCalc{}

	sum, err := calc.Add("//#\n1#2")
	failOnError(err)

	fmt.Println(sum)
}

func failOnError(err error) {
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
}
