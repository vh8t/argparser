package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vh8t/argparser"
)

func main() {
	parser := argparser.NewRule("greet", "Greet a person", "v1.0", true).
		AddStringFlag("name", "n", "Name of the person to greet", true).
		AddBoolFlag("shout", "", "Whether to shout the greeting")

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

	name, _ := parser.GetStringFlag("name")

	greeting := fmt.Sprintf("Hello, %s!", name)
	if parser.GetBoolFlag("shout") {
		greeting = strings.ToUpper(greeting)
	}

	fmt.Println(greeting)
}
