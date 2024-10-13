package flax

import (
	"iter"
	"os"
	"strings"
)
//go:generate stringer -type=ArgType
type ArgType uint8

const (
	ArgTypeShort ArgType = iota
	ArgTypeLong
	ArgTypeDoubleDash
	ArgTypeProgram
	ArgTypeInvalid
)

type arg struct {
	Type  ArgType
	Name string
	Raw  string
}

func parseArg(raw string, n uint) *arg {
	name := raw
	prefixCount := 0
	for {
		rest, found := strings.CutPrefix(name, "-")
		if !found || prefixCount >= len("--") {
			break
		}
		prefixCount++
		name = rest
	}
	typ := ArgTypeInvalid
	if raw == "--" {
		typ = ArgTypeDoubleDash
	} else if prefixCount == 1 {
		typ = ArgTypeShort
	} else if prefixCount == 2 {
		typ = ArgTypeLong
	} else if n == 0 {
		typ = ArgTypeProgram
	}
	return &arg{typ, name, raw}
}

type Shifter struct {
	args []string
	argc uint
	n    uint
}

var shifter = Shifter{os.Args, uint(len(os.Args)), 0}

func (s *Shifter) shifted() iter.Seq[*arg] {
	return func(yield func(*arg) bool) {
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

func Shift() *arg {
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

func Peek() *arg {
	arg := Shift()
	Unshift()
	return arg
}

func Collect() (program *arg, rest []*arg) {
	program = Shift()
	for Peek() != nil {
		arg := Shift()
		rest = append(rest, arg)
	}
	return program, rest
}
