package flax

import (
	"errors"
)

type FlagSpec struct {
	Name     string
	Short    string
	Help     string
	Required bool
}

type flag struct {
	// serial is the readonly value from the flag parser.
	// it should be used to parse to typed values and never be mutated.
	serial string
	spec   *FlagSpec
}

type CLI struct {
	flags map[string]flag
}

func NewCLI() *CLI {
	return &CLI{
		flags: map[string]flag{},
	}
}

var serialUndef = "UNDEF"

var errFlagNotFound = errors.New("flag not found")

func (c *CLI) Set(spec *FlagSpec) *CLI {
	c.flags[spec.Name] = flag{serialUndef, spec}
	return c
}

func (c *CLI) Get(name string) *flag {
	if flag, ok := c.flags[name]; ok {
		return &flag
	}
	return nil
}

func (c *CLI) Parse(args []string) {
	todo("implement parse")
}
