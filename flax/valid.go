package flax

import (
	"fmt"
	"strconv"
)

// TODO: float, uint, 32 bit, boolean, comma,space,tab separated

type intValidator struct {
	f *func(v int) error
}

func Int() *intValidator                        { return &intValidator{} }
func IntFunc(f func(v int) error) *intValidator { return &intValidator{&f} }

func (v *intValidator) toInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	value := int(i)
	if err != nil {
		return value, err
	}
	return value, nil
}

// TODO: TryFunc would be better I think
func (v *intValidator) Try(f *flag) (int, error) {
	if f.serial == Undef {
		return 0, fmt.Errorf("`%s` expects an argument", f.spec.Name)
	}
	value, err := v.toInt(f.serial)
	if err != nil {
		return value, err
	}
	if v.f != nil {
		if err := (*v.f)(value); err != nil {
			return value, err
		}
	}
	return value, err
}

// type fileValidator struct {
// 	f *func(v *os.File) error
// }
// func Dir() *fileValidator                             { return &fileValidator{} }
// func DirFunc(f func(v *os.File) error) *fileValidator { return &fileValidator{&f} }
//
// func (v *fileValidator) toFile(s string) (os.File, error) {
// 	fifo, err := os.Stat(s)
//     fifi
// 	return *file, err
// }
//
// func (v *fileValidator) Try(f *flag) (*os.File, error) {
// 	value, err := v.toFile(f.serial)
// 	if err != nil {
// 		return value, err
// 	}
// 	if v.f != nil {
// 		if err := (*v.f)(value); err != nil {
// 			return value, err
// 		}
// 	}
// 	return value, err
// }

type stringValidator struct {
	f *func(v string) error
}

func String() *stringValidator                           { return &stringValidator{} }
func StringFunc(f func(v string) error) *stringValidator { return &stringValidator{&f} }

func (v *stringValidator) toString(s string) (string, error) { return s, nil }
func (v *stringValidator) Try(f *flag) (string, error) {
	if f.serial == Undef {
		return "", fmt.Errorf("`%s` expects an argument", f.spec.Name)
	}
	value, err := v.toString(f.serial)
	if err != nil {
		return value, err
	}
	if v.f != nil {
		if err := (*v.f)(value); err != nil {
			return value, err
		}
	}
	return value, err
}
