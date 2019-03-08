package main

import (
	"fmt"

	"github.com/screwyprof/stringcalc"
)

func main() {
	calc := stringcalc.StringCalc{}
	sum := calc.Add("5")

	fmt.Println(sum)
}
