package main

import (
	"argparser"
	"fmt"
	"log"
	"os"
)

func main() {
	rule := argparser.NewRule("demo", "Argparser demo", "v2.2", true).
		AddPositional("name").
		AddBoolFlag("bold", "", "use bold format for text").
		AddBoolFlag("italic", "", "use italic format for text").
		AddIntFlag("red", "r", "red value for text color").
		AddIntFlag("green", "g", "green value for text color").
		AddIntFlag("blue", "b", "blue value for text color")

	err := rule.Parse(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	name := rule.GetPositional("name")

	bold := rule.GetBoolFlag("bold")
	italic := rule.GetBoolFlag("italic")

	red, rOk := rule.GetIntFlag("red")
	green, gOk := rule.GetIntFlag("green")
	blue, bOk := rule.GetIntFlag("blue")

	var format string

	if bold {
		format += "1"
	}

	if italic {
		if len(format) != 0 {
			format += ";"
		}
		format += "3"
	}

	if rOk && gOk && bOk {
		if len(format) != 0 {
			format += ";"
		}
		format += fmt.Sprintf("38;2;%d;%d;%dm", red, green, blue)
	}

	if len(format) != 0 {
		format = "\x1b[" + format
	}

	fmt.Printf("%s%s\n", format, name)
}
