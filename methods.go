package argparser

import "fmt"

func (r *Rule) AddStringFlag(long, short, description string) *Rule {
	r.stringFlags = append(r.stringFlags, StringFlag{
		long:        long,
		short:       short,
		description: description,
		empty:       true,
		value:       "",
	})

	return r
}

func (r *Rule) GetStringFlag(name string) (string, bool) {
	for _, flag := range r.stringFlags {
		if flag.long == name {
			return flag.value, !flag.empty
		}
	}

	return "", false
}

func (r *Rule) AddIntFlag(long, short, description string) *Rule {
	r.intFlags = append(r.intFlags, IntFlag{
		long:        long,
		short:       short,
		description: description,
		empty:       true,
		value:       0,
	})

	return r
}

func (r *Rule) GetIntFlag(name string) (int, bool) {
	for _, flag := range r.intFlags {
		if flag.long == name {
			return flag.value, !flag.empty
		}
	}

	return 0, false
}

func (r *Rule) AddFloatFlag(long, short, description string) *Rule {
	r.floatFlags = append(r.floatFlags, FloatFlag{
		long:        long,
		short:       short,
		description: description,
		empty:       true,
		value:       0.0,
	})

	return r
}

func (r *Rule) GetFloatFlag(name string) (float64, bool) {
	for _, flag := range r.floatFlags {
		if flag.long == name {
			return flag.value, !flag.empty
		}
	}

	return 0, false
}

func (r *Rule) AddBoolFlag(long, short, description string) *Rule {
	r.boolFlags = append(r.boolFlags, BoolFlag{
		long:        long,
		short:       short,
		description: description,
		empty:       true,
		value:       false,
	})

	return r
}

func (r *Rule) GetBoolFlag(name string) bool {
	for _, flag := range r.boolFlags {
		if flag.long == name {
			return flag.value
		}
	}

	return false
}

func (r *Rule) AddPositional(name string) *Rule {
	r.positionals = append(r.positionals, Positional{
		name:  name,
		empty: true,
		value: "",
	})

	return r
}

func (r *Rule) GetPositional(name string) string {
	for _, flag := range r.positionals {
		if flag.name == name {
			return flag.value
		}
	}

	return ""
}

func (r *Rule) Help() string {
	out := fmt.Sprintf("%s %s\n\nUSAGE:\n  %s", r.program, r.version, r.program)

	var maxLBool, maxSBool int
	var maxLArg, maxSArg int

	for _, flag := range r.boolFlags {
		if len(flag.long)+2 > maxLBool && len(flag.long) != 0 {
			maxLBool = len(flag.long) + 2
		}

		if len(flag.short)+1 > maxSBool && len(flag.short) != 0 {
			maxSBool = len(flag.short) + 1
		}
	}

	for _, flag := range r.stringFlags {
		if len(flag.long)+2 > maxLArg && len(flag.long) != 0 {
			maxLArg = len(flag.long) + 2
		}

		if len(flag.short)+1 > maxSArg && len(flag.short) != 0 {
			maxSArg = len(flag.short) + 1
		}
	}

	for _, flag := range r.intFlags {
		if len(flag.long)+2 > maxLArg && len(flag.long) != 0 {
			maxLArg = len(flag.long) + 2
		}

		if len(flag.short)+1 > maxSArg && len(flag.short) != 0 {
			maxSArg = len(flag.short) + 1
		}
	}

	if len(r.boolFlags) > 0 {
		out += " [FLAGS]"
	}

	if len(r.stringFlags) > 0 || len(r.intFlags) > 0 {
		out += " [OPTIONS]"
	}

	for _, pos := range r.positionals {
		out += fmt.Sprintf(" <%s>", pos.name)
	}

	if len(r.boolFlags) > 0 {
		out += "\n\nFLAGS:"
		for _, flag := range r.boolFlags {
			var short, long string
			if len(flag.short) > 0 {
				short = fmt.Sprintf("-%s", flag.short)
			}
			if len(flag.long) > 0 {
				long = fmt.Sprintf("--%s", flag.long)
			}

			out += fmt.Sprintf("\n  %-*s %-*s   %s", maxSBool, short, maxLBool, long, flag.description)
		}
	}

	if len(r.stringFlags) > 0 || len(r.intFlags) > 0 {
		out += "\n\nOPTIONS:"
		for _, flag := range r.stringFlags {
			var short, long string
			if len(flag.short) > 0 {
				short = fmt.Sprintf("-%s", flag.short)
			}
			if len(flag.long) > 0 {
				long = fmt.Sprintf("--%s", flag.long)
			}

			out += fmt.Sprintf("\n  %-*s %-*s <str>   %s", maxSArg, short, maxLArg, long, flag.description)
		}

		for _, flag := range r.intFlags {
			var short, long string
			if len(flag.short) > 0 {
				short = fmt.Sprintf("-%s", flag.short)
			}
			if len(flag.long) > 0 {
				long = fmt.Sprintf("--%s", flag.long)
			}

			out += fmt.Sprintf("\n  %-*s %-*s <int>   %s", maxSArg, short, maxLArg, long, flag.description)
		}

		for _, flag := range r.floatFlags {
			var short, long string
			if len(flag.short) > 0 {
				short = fmt.Sprintf("-%s", flag.short)
			}
			if len(flag.long) > 0 {
				long = fmt.Sprintf("--%s", flag.long)
			}

			out += fmt.Sprintf("\n  %-*s %-*s <flt>   %s", maxSArg, short, maxLArg, long, flag.description)
		}
	}

	return out
}
