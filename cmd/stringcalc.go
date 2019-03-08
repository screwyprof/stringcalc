package main

import (
	"fmt"

	"github.com/screwyprof/stringcalc"
)

func main() {
	calc := stringcalc.StringCalc{}
	sum := calc.Add("1\n2,3")

	fmt.Println(sum)
}
