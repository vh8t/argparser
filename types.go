package argparser

type Rule struct {
	program     string
	description string
	version     string
	helpCommand bool
	boolFlags   []BoolFlag
	stringFlags []StringFlag
	intFlags    []IntFlag
	floatFlags  []FloatFlag
	positionals []Positional
}

type BoolFlag struct {
	long        string
	short       string
	description string
	empty       bool
	value       bool
}

type StringFlag struct {
	long        string
	short       string
	description string
	required    bool
	empty       bool
	value       string
}

type IntFlag struct {
	long        string
	short       string
	description string
	required    bool
	empty       bool
	value       int
}

type FloatFlag struct {
	long        string
	short       string
	description string
	required    bool
	empty       bool
	value       float64
}

type Positional struct {
	name  string
	empty bool
	value string
}
