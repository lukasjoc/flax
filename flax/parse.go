package flax

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type argKind int

const (
	argKindLong argKind = iota
	argKindShort
)

var (
	longArgPrefix     = "--"
	shortArgPrefix    = "-"
	longArgPrefixLen  = len(longArgPrefix)
	shortArgPrefixLen = len(shortArgPrefix)
)

//	type arg struct {
//		kind  argKind
//		name  string
//		value []string
//	}
type parser struct {
	c    *cmd
	rest []string
	// parsed []arg
}

func (p *parser) shift()        { p.rest = p.rest[1:] }
func (p *parser) shifted() bool { return len(p.rest) == 0 }
func (p *parser) head() string {
	if p.shifted() {
		return ""
	}
	return p.rest[0]
}
func (p *parser) next() (string, error) {
	if p.shifted() {
		return "", errors.New("end of input")
	}
	temp := p.head()
	p.shift()
	return temp, nil
}
func (p *parser) nextIsLong() bool {
	return strings.HasPrefix(p.head(), longArgPrefix)
}
func (p *parser) nextIsShort() bool {
	return !p.nextIsLong() && strings.HasPrefix(p.head(), shortArgPrefix)
}
func (p *parser) eatWhileValue() []string {
	s := []string{}
	for {
		if p.shifted() || p.nextIsLong() || p.nextIsShort() {
			break
		}
		a, err := p.next()
		unreachable(err != nil, "we already check this above")
		s = append(s, a)
	}
	return s
}
func (p *parser) eatWhileLong() error {
	if !p.nextIsLong() {
		return nil
	}
	a, err := p.next()
	if err != nil {
		return err
	}
	// TODO: validate flag ident
	if len(a) <= longArgPrefixLen {
		return fmt.Errorf("cannot parse: `%s`", a)
	}
	name := a[longArgPrefixLen:]
	flag := p.c.Flag(name)
	if flag == nil {
		return fmt.Errorf("flag not found: `%s`", name)
	}
	value := p.eatWhileValue()
	if len(value) > 1 {
		return fmt.Errorf("multi-value arg not supported: `%s`", name)
	}
	// TODO: type casting
	// TODO: validators
	flag.setValue(value[0])
	return nil
}
func (p *parser) eatWhileShort() error {
	if !p.nextIsLong() {
		return nil
	}
	// a, err := p.next()
	// if err != nil {
	// 	return
	// }
	// // TOearO: validate flag ident
	// if len(a) <= shortArgPrefixLen {
	// 	panic(fmt.Sprintf("invalid short arg: %v", a))
	// }
	// p.parsed = append(p.parsed, arg{argKindLong, a[shortArgPrefixLen:], p.eatWhileValue()})
	return nil
}

func (p *parser) parse() error {
	for !p.shifted() {
		if err := p.eatWhileLong(); err != nil {
			return err
			// } else if err := p.eatWhileShort(); err != nil {
			// 	return err
		} else {
			fmt.Errorf("cannot parse: %v", p.head())
		}
	}
	return nil
}

func Parse(c *cmd) error {
	p := parser{c, os.Args[1:]}
	// TODO: sanitise the args provided ~before~ calling parse
	return p.parse()
}
