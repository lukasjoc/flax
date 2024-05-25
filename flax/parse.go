package flax

import (
	"errors"
	"fmt"
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

type arg struct {
	kind  argKind
	name  string
	value string
}

type parser struct {
	rest []string
	args []arg
}

func newParser(args []string) *parser {
	return &parser{args[1:], []arg{}}
}

func (p *parser) shift()        { p.rest = p.rest[1:] }
func (p *parser) shifted() bool { return len(p.rest) == 0 }
func (p *parser) front() string { return p.rest[0] }
func (p *parser) next() (string, error) {
	if p.shifted() {
		return "", errors.New("end of input")
	}
	temp := p.front()
	p.shift()
	return temp, nil
}
func (p *parser) nextIsLong() bool {
	if p.shifted() {
		return false
	}
	return strings.HasPrefix(p.front(), longArgPrefix)
}
func (p *parser) nextIsShort() bool {
	if p.shifted() || p.nextIsLong() {
		return false
	}
	return strings.HasPrefix(p.front(), shortArgPrefix)
}
func (p *parser) eatWhileValue() []string {
	s := []string{}
	for {
		if p.shifted() || p.nextIsLong() || p.nextIsShort() {
			break
		}
		a, err := p.next()
		unreachable(err != nil, "we should never have and err here")
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
	value := p.eatWhileValue()
	if len(value) > 1 {
		return fmt.Errorf("cannot parse multi-value arg: `%s`", name)
	}
	if len(value) > 0 {
		p.args = append(p.args, arg{argKindLong, name, value[0]})
	} else {
		p.args = append(p.args, arg{argKindLong, name, Undef})
	}
	return nil
}

func (p *parser) eatWhileShort() error {
	if !p.nextIsShort() {
		return nil
	}
	a, err := p.next()
	if err != nil {
		return err
	}
	// TODO: validate flag ident
	if len(a) <= shortArgPrefixLen {
		return fmt.Errorf("invalid flag syntax or value: `%s`", a)
	}
	name := a[shortArgPrefixLen:]
	value := p.eatWhileValue()
	if len(value) > 1 {
		return fmt.Errorf("cannot parse multi-value arg: `%s`", name)
	}
	if len(value) > 0 {
		p.args = append(p.args, arg{argKindShort, name, value[0]})
	} else {
		p.args = append(p.args, arg{argKindShort, name, Undef})
	}
	return nil
}

func (p *parser) parse() error {
	for !p.shifted() {
		a := p.front()
		switch {
		case p.nextIsLong():
			if err := p.eatWhileLong(); err != nil {
				return err
			}
		case p.nextIsShort():
			if err := p.eatWhileShort(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid flag syntax or value: `%s`", a)
		}
	}
	return nil
}
