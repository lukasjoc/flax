package flax

import (
	"errors"
	"fmt"
	"os"
	"slices"
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
	Name  string
	flags map[string]*flag
}

func NewCLI(name string) *CLI {
	return &CLI{
		Name:  name,
		flags: map[string]*flag{},
	}
}

var errFlagNotFound = errors.New("flag not found")

func (c *CLI) Set(spec *FlagSpec) *CLI {
	// TODO: bail if flag name already exists (short, and long)
	c.flags[spec.Name] = &flag{Undef, spec}
	return c
}

func (c *CLI) Get(name string) *flag {
	if flag, ok := c.flags[name]; ok {
		return flag
	}
	return nil
}

func (c *CLI) Parse(args []string) error {
	p := newParser(args)
	if err := p.parse(); err != nil {
		return err
	}
	for _, flag := range c.flags {
		f := flag
		okidx := slices.IndexFunc(p.args, func(a arg) bool {
			return flag.spec.Name == a.name ||
				flag.spec.Short == a.name
		})
		if flag.spec.Required && okidx == -1 {
			return fmt.Errorf("missing requried flag `%s`", flag.spec.Name)
		}
		if okidx == -1 {
			continue
		}
		unreachable(okidx >= len(p.args), "okidx should always be in range")
		f.serial = p.args[okidx].value
	}
	return nil
}

func Bail(err error) {
	fmt.Fprintf(os.Stderr, "Usage: TODO\n")
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}
