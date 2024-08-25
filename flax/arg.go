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
	Name       string
	Raw        string
	Short      bool
	Long       bool
	Program    bool
	DoubleDash bool
	// TODO: pack the arguments to short and long flags into the argument
	// until a different flag is reached. Other args are treated as positionals.
	// Values  []struct{Raw: string, Type: func(a string) T {/*infer the type through reflection*/} }
	// TODO: Support for clobbered args -ArGs. Those should be parsed as seperate short args.
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
		} else {
			prefixCount++
		}
		name = rest
	}
	return &Arg{name,
		raw,
		prefixCount == 1,
		prefixCount == 2,
		n == 0,
		raw == "--",
	}
}

type Item *Arg

func (s *Shifter) shifted() iter.Seq[Item] {
	return func(yield func(Item) bool) {
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

func Shift() Item {
	next, stop := iter.Pull(shifter.shifted())
	defer stop()
	arg, ok := next()
	if !ok {
		return nil
	}
	return arg
}

func Unshift() Item {
	panic("TODO: not implemented yet !")
	return nil
}

func Peek() Item {
	panic("TODO: not implemented yet !")
	return nil
}
