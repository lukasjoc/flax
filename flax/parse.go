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
	// argKindValue
	// argKindCmd
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
	value []string
}
type parser struct {
	rest   []string
	parsed []arg
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
func (p *parser) eatValues() []string {
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
func (p *parser) eatWhileLong() {
	a, err := p.next()
	if err != nil {
		return
	}
	// TODO: validate flag ident
	if len(a) <= longArgPrefixLen {
		panic(fmt.Sprintf("invalid long arg: %v", a))
	}
	p.parsed = append(p.parsed, arg{argKindLong, a[longArgPrefixLen:], p.eatValues()})
}
func (p *parser) eatWhileShort() {
	a, err := p.next()
	if err != nil {
		return
	}
	// TOearO: validate flag ident
	if len(a) <= shortArgPrefixLen {
		panic(fmt.Sprintf("invalid short arg: %v", a))
	}
	p.parsed = append(p.parsed, arg{argKindLong, a[shortArgPrefixLen:], p.eatValues()})
}
func (p *parser) parse() {
	for !p.shifted() {
		if p.nextIsLong() {
			p.eatWhileLong()
		} else if p.nextIsShort() {
			p.eatWhileShort()
		} else {
			panic(fmt.Sprintf("cannot parse: %v", p.head()))
		}
	}
}

func Parse(c *cmd) {
	p := parser{os.Args[1:], []arg{}}
	// TODO: sanitise the args provided ~before~ calling parse
	p.parse()
	fmt.Printf("parsed args: %#+v \n", p.parsed)
}
