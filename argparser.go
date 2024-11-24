package argparser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func NewRule(program, description, version string, helpCommand bool) *Rule {
	rule := Rule{
		program:     program,
		description: description,
		version:     version,
		helpCommand: helpCommand,
	}

	if helpCommand {
		rule.boolFlags = append(rule.boolFlags, BoolFlag{
			long:        "help",
			short:       "",
			description: "show this message",
			empty:       true,
			value:       false,
		})
	}

	return &rule
}

func clean[T any](list []T, indexes []int) []T {
	var cleaned []T
	sort.Slice(indexes, func(i, j int) bool {
		return j < i
	})

	for i, val := range list {
		allowed := true
		for _, idx := range indexes {
			if i == idx {
				allowed = false
			}
		}
		if allowed {
			cleaned = append(cleaned, val)
		}
	}

	return cleaned
}

func (r *Rule) Parse(args []string) error {
	if r.helpCommand {
		for _, arg := range args {
			if arg == "--help" {
				fmt.Println(r.Help())
				return nil
			}
		}
	}

	var remove []int

	for i, flag := range r.boolFlags {
		for j, arg := range args {
			if (strings.HasPrefix(arg, "--") && strings.TrimPrefix(arg, "--") == flag.long) || (strings.HasPrefix(arg, "-") && strings.TrimPrefix(arg, "-") == flag.short) {
				if !r.boolFlags[i].empty {
					return fmt.Errorf("%s is duplicate flag", arg)
				}

				r.boolFlags[i].empty = false
				r.boolFlags[i].value = true
				remove = append(remove, j)
			}
		}
	}

	args = clean(args, remove)
	remove = []int{}

	for i, flag := range r.stringFlags {
		for j, arg := range args {
			if (strings.HasPrefix(arg, "--") && strings.TrimPrefix(arg, "--") == flag.long) || (strings.HasPrefix(arg, "-") && strings.TrimPrefix(arg, "-") == flag.short) {
				if !r.stringFlags[i].empty {
					return fmt.Errorf("%s is duplicate flag", arg)
				}

				if j+1 >= len(args) {
					return fmt.Errorf("%s flag is missing argument", arg)
				}

				r.stringFlags[i].empty = false
				r.stringFlags[i].value = args[j+1]
				remove = append(remove, j, j+1)
			}
		}
	}

	args = clean(args, remove)
	remove = []int{}

	for i, flag := range r.intFlags {
		for j, arg := range args {
			if (strings.HasPrefix(arg, "--") && strings.TrimPrefix(arg, "--") == flag.long) || (strings.HasPrefix(arg, "-") && strings.TrimPrefix(arg, "-") == flag.short) {
				if !r.intFlags[i].empty {
					return fmt.Errorf("%s is duplicate flag", arg)
				}

				if j+1 >= len(args) {
					return fmt.Errorf("%s flag is missing argument", arg)
				}

				value, err := strconv.ParseInt(args[j+1], 0, 64)
				if err != nil {
					return err
				}

				r.intFlags[i].empty = false
				r.intFlags[i].value = int(value)
				remove = append(remove, j, j+1)
			}
		}
	}

	args = clean(args, remove)
	remove = []int{}

	for i, flag := range r.floatFlags {
		for j, arg := range args {
			if (strings.HasPrefix(arg, "--") && strings.TrimPrefix(arg, "--") == flag.long) || (strings.HasPrefix(arg, "-") && strings.TrimPrefix(arg, "-") == flag.short) {
				if !r.intFlags[i].empty {
					return fmt.Errorf("%s is duplicate flag", arg)
				}

				if j+1 >= len(args) {
					return fmt.Errorf("%s flag is missing argument", arg)
				}

				value, err := strconv.ParseFloat(args[j+1], 64)
				if err != nil {
					return err
				}

				r.floatFlags[i].empty = false
				r.floatFlags[i].value = value
				remove = append(remove, j, j+1)
			}
		}
	}

	args = clean(args, remove)
	remove = []int{}

	for i, pos := range r.positionals {
		if i >= len(args) {
			return fmt.Errorf("Missing required field <%s>", pos.name)
		}

		r.positionals[i].empty = false
		r.positionals[i].value = args[i]
	}

	return nil
}
