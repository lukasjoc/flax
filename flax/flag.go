package flax

import (
	"errors"
	"reflect"
	"strings"
)

func NewCmd(name string) *cmd {
	return &cmd{name, map[string]*flag[any]{}}
}

type cmd struct {
	name  string
	flags map[string]*flag[any]
}

func (c *cmd) NewFlag(name string) *flag[any] {
	// TODO: not ideal but generics still bad
	return c.setFlag(&flag[any]{cmd: c, name: name})
}

func (c *cmd) setFlag(f *flag[any]) *flag[any] {
	if _, ok := c.flags[f.name]; ok {
		panic("flag with that name already exists")
	}
	c.flags[f.name] = f
	return f
}

func (c *cmd) Flag(name string) *flag[any] {
	if flag, ok := c.flags[name]; ok {
		return flag
	}
	panic("flag with that name doesnt exists")
}

type flag[T any] struct {
	cmd        *cmd
	name       string
	short      string
	help       string
	required   bool
	value      *T
	validators []func(v T) error
}

func (f *flag[T]) Cmd() *cmd { return f.cmd }

func (f *flag[T]) Short() *flag[T] {
	if len(f.name) == 0 {
		return nil /*should never branch into here */
	}
	f.short = string(strings.ToLower(f.name)[0])
	return f
}
func (f *flag[T]) Help(help string) *flag[T] {
	f.help = help
	return f
}

func (f *flag[T]) Required() *flag[T] {
	f.required = true
	return f
}

func (f *flag[T]) Default(v T) *flag[T] {
	f.value = &v
	return f
}

func (f *flag[T]) Value() *T {
	return f.value
}

func (f *flag[T]) setValue(v T) {
	f.value = &v
}

func (f *flag[T]) ValueOr(or T) any {
	if f.value == nil {
		return or
	}
	return f.value
}

func (f *flag[T]) Validator(validatorFunc func(v T) error) *flag[T] {
	f.validators = append(f.validators, validatorFunc)
	return f
}

// TODO: maybe later convert to a single type validator: func Type(reflec.Kind) error
func (f *flag[T]) String() *flag[T] {
	validatorFunc := func(v T) error {
		if reflect.TypeOf(f.value).Kind() != reflect.String {
			return errors.New("not a valid string type")
		}
		return nil
	}
	f.validators = append(f.validators, validatorFunc)
	return f
}

func (f *flag[T]) Int() *flag[T] {
	validatorFunc := func(v T) error {
		if reflect.TypeOf(f.value).Kind() != reflect.Int ||
			reflect.TypeOf(f.value).Kind() != reflect.Uint {
			return errors.New("not a valid int type")
		}
		return nil
	}
	f.validators = append(f.validators, validatorFunc)
	return f
}

func (c *cmd) Parse() { Parse(c) }
