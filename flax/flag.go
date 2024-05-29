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
	flags []*flag
	Name  string
	Desc  string
}

func NewCLI(name string, desc string) *CLI {
	cli := CLI{
		flags: []*flag{},
		Name:  name,
		Desc:  desc,
	}
	return &cli
}

var errFlagNotFound = errors.New("flag not found")

func (c *CLI) Set(spec *FlagSpec) *CLI {
	if len(spec.Name) == 0 && len(spec.Short) == 0 {
		panic("flag neeeds a name or a short name")
	}
	// TODO: bail if flag name already exists (short, and long)
	c.flags = append(c.flags, &flag{Undef, spec})
	return c
}

func (c *CLI) Get(name string) *flag {
	index := slices.IndexFunc(c.flags, func(f *flag) bool {
		return f.spec.Name == name
	})
	if index == -1 {
		return nil
	}
	return c.flags[index]
}

func (c *CLI) Parse(args []string) error {
	p := newParser(args)
	if err := p.parse(); err != nil {
		return err
	}
	var err error
	for _, flag := range c.flags {
		f := flag
		okidx := slices.IndexFunc(p.args, func(a arg) bool {
			return flag.spec.Name == a.name || flag.spec.Short == a.name
		})
		if flag.spec.Required && okidx == -1 {
			if len(flag.spec.Name) > 0 {
				err = fmt.Errorf("missing requried flag: `%s`", flag.spec.Name)
			}
			continue
		}
		if okidx == -1 {
			continue
		}
		unreachable(okidx >= len(p.args), "okidx should always be in range")
		f.serial = p.args[okidx].value
	}

	c.Set(&FlagSpec{Name: "help", Short: "h", Help: "Show this help."})
	help := slices.IndexFunc(p.args, func(a arg) bool {
		return a.name == "help" || a.name == "h"
	})
	if help != -1 {
		c.Exit(nil)
	}
	return err
}

func (c *CLI) PrintHelp() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flag]\n", c.Name)
	fmt.Fprintf(os.Stderr, "\nDescription:\n%s\n", c.Desc)
	fmt.Fprintf(os.Stderr, "\nFlags:\n")
	for _, flag := range c.flags {
		s := ""
		if len(flag.spec.Short) > 0 {
			s += fmt.Sprintf("-%s", flag.spec.Short)
		}
		if len(flag.spec.Name) > 0 {
			if len(flag.spec.Short) != 0 {
                s += fmt.Sprintf(",")
			}
			s += fmt.Sprintf("--%s", flag.spec.Name)
		}
		if len(flag.spec.Help) > 0 {
			s += fmt.Sprintf("\n%s", flag.spec.Help)
		}
        if flag.spec.Required {
			s += fmt.Sprintf("(\033[1;31m*\033[0;m required)")
		}
		fmt.Fprintf(os.Stderr, "%s\n", s)
	}
}
func (c *CLI) Exit(err error) {
	c.PrintHelp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n\033[1;31mERROR:\033[0;m %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
