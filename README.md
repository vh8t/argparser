# argparser

Simple argument parser for go

## Creating a parser

To create a parser there is a methos `NewRule` which takes in 4 arguments

1. `program` (`string`) - program name
1. `description` (`string`) - program description
1. `version` (`string`) - program version
1. `helpCommand` (`bool`) - have the help flag by default

## Setting up some rules

Parser has many methods to add rules, all the flag names must be without the `--` and `-` prefix

### Setting and arguments

1. `AddStringFlag`
    - `long` (`string`) - long flag name, it is required to have this value
    - `short` (`string`) - short flag name
    - `description` (`string`) - description of the flag
    - `required` (`bool`) - if the flag is required or not
    - `defaultValue` (optional `string`) - default value for the flag, works only if the `required` is `false`

1. `AddIntFlag`
    - `long` (`string`) - long flag name, it is required to have this value
    - `short` (`string`) - short flag name
    - `description` (`string`) - description of the flag
    - `required` (`bool`) - if the flag is required or not
    - `defaultValue` (optional `int`) - default value for the flag, works only if the `required` is `false`

1. `AddFloatFlag`
    - `long` (`string`) - long flag name, it is required to have this value
    - `short` (`string`) - short flag name
    - `description` (`string`) - description of the flag
    - `required` (`bool`) - if the flag is required or not
    - `defaultValue` (optional `float64`) - default value for the flag, works only if the `required` is `false`

1. `AddBoolFlag`
    - `long` (`string`) - long flag name, it is required to have this value
    - `short` (`string`) - short flag name
    - `description` (`string`) - description of the flag

1. `AddPositional`
    - `name` (`string`) - flag name, it is required to have this value

## Parsing

Once you have set up your rules you can call the `Parse` method on the rule, it takes in 1 argument

- `args` (`[]string`) - these are the arguments to parse, if using `os.Args` don't include the first argument

## Getting the values

To get the values you can use one of many getter methods on the rule after parsing, all of them have 1 same argument

- `name` (`string`) - this is the name or long value you used for the flags/positionals

### Getters and return values

1. `GetStringFlag`
    - `string` - the value
    - `bool` - ok value, `false` if value wasnt set

1. `GetIntFlag`
    - `int` - the value
    - `bool` - ok value, `false` if value wasnt set

1. `GetFloatFlag`
    - `float64` - the value
    - `bool` - ok value, `false` if value wasnt set

1. `GetBoolFlag`
    - `bool` - the value

1. `GetPositional`
    - `string` - the value
