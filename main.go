package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	cli := flax.NewCLI("foobar")
	cli.Set(&flax.FlagSpec{Name: "firstname", Short: "f", Required: true, Help: "provide your firstname"})
	cli.Set(&flax.FlagSpec{Name: "lastname", Short: "l", Required: true, Help: "provide your lastname"})
	cli.Set(&flax.FlagSpec{Name: "profession", Short: "p", Help: "provide your profession"})

	if err := cli.Parse(os.Args); err != nil {
		flax.Bail(err)
	}

	firstname, err := flax.String().Try(cli.Get("firstname"))
	if err != nil {
		flax.Bail(err)
	}
	fmt.Println("My firstname is: ", firstname)

	lastname, err := flax.String().Try(cli.Get("lastname"))
	if err != nil {
		flax.Bail(err)
	}
	fmt.Println("My lastname is: ", lastname)

	profession, err := flax.StringFunc(func(v string) error {
		if strings.Contains(v, "software") {
			return errors.New("touch some grass")
		}
		return nil
	}).Try(cli.Get("profession"))
	if err != nil {
		flax.Bail(err)
	}
	fmt.Println("My profession is: ", profession)
}
