package main

import (
	"fmt"
	"github.com/screwyprof/stringcalc"
)

func main() {
	run()
}

func run() {
	calc := stringcalc.StringCalc{}

	data := []string{"", "5", "5,2", "3,2,1,0,1", "1\n2,3", "non-numbers",
		"//;\n1;2", "1,-2,-4", "2,1001", "//[***]\n1***2***3", "//[*][%]\n1*2%3"}
	for _, input := range data {
		addAndPrintSum(calc, input)
	}
}

func addAndPrintSum(calc stringcalc.StringCalc, input string) {
	sum, err := calc.Add(input)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}
	fmt.Println(sum)
}
