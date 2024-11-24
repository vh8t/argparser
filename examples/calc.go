package main

import (
	"fmt"
	"os"

	"github.com/vh8t/argparser"
)

func main() {
	parser := argparser.NewRule("calc", "Perform binary calculation", "v1.0", true).
		AddIntFlag("num1", "", "Number 1", true).
		AddIntFlag("num2", "", "Number 2", true).
		AddPositional("operation")

	err := parser.Parse(os.Args[1:])
	help := parser.GetBoolFlag("help")

	if err != nil && !help {
		fmt.Println(err)
		fmt.Println(parser.Help())
		return
	} else if help {
		fmt.Println(parser.Help())
		return
	}

	num1, _ := parser.GetIntFlag("num1")
	num2, _ := parser.GetIntFlag("num2")

	switch parser.GetPositional("operation") {
	case "add", "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "sub", "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "mul", "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "div", "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	default:
		fmt.Println("operation can only be one of following:\n  add, sub, mul, div")
		fmt.Println(parser.Help())
	}
}
