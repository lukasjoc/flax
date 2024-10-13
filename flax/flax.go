package flax

import (
	"iter"
	"os"
	"strings"
)

type Shifter struct {
	args []string
	argc uint
	n    uint
}

var shifter = Shifter{os.Args, uint(len(os.Args)), 0}

type Arg struct {
	// The name of the arg. Without dashes if its a flag.
	Name string
	// The raw version for the arg. With dashes if its a flag.
	Raw string
	// Whether its a short arg. E.g. `-short`
	Short bool
	// Wether its a long arg. E.g. `--long`
	Long bool
	// Whether its the program/executable.
	Program bool
	// Whether its a double dash arg E.g. `--`
	DoubleDash bool
}

func parseArg(raw string, n uint) *Arg {
	const (
		flagPrefix     = "-"
		maxPrefixCount = 2
	)
	name := raw
	prefixCount := 0
	for {
		rest, found := strings.CutPrefix(name, flagPrefix)
		if !found || prefixCount >= maxPrefixCount {
			break
		}
		prefixCount++
		name = rest
	}
	return &Arg{
		Name:       name,
		Raw:        raw,
		Short:      prefixCount == 1,
		Long:       prefixCount == 2,
		Program:    n == 0,
		DoubleDash: raw == "--",
	}
}

func (s *Shifter) shifted() iter.Seq[*Arg] {
	return func(yield func(*Arg) bool) {
		if s.n+1 > s.argc {
			return
		}
		arg := parseArg(s.args[s.n], s.n)
		s.n++
		if !yield(arg) {
			return
		}
	}
}

func Shift() *Arg {
	next, stop := iter.Pull(shifter.shifted())
	defer stop()
	arg, ok := next()
	if !ok {
		return nil
	}
	return arg
}

func Unshift() uint {
	if shifter.n > 0 {
		shifter.n--
	}
	return shifter.n
}

func Peek() *Arg {
	arg := Shift()
	Unshift()
	return arg
}

func Collect() (program *Arg, rest []*Arg) {
	program = Shift()
	for Peek() != nil {
		arg := Shift()
		rest = append(rest, arg)
	}
	return program, rest
}

// TODO: pack the arguments to short and long flags into the argument
// until a different flag is reached. Other args are treated as positionals.
// Values  []struct{Raw: string, Type: func(a string) T {/*infer the type through reflection*/} }
// TODO: Support for clobbered args -ArGs. Those should be parsed as seperate short args.
// TODO: `-` should be disallowed + more argument text validation.
