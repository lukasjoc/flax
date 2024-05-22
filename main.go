package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	cmd := flax.NewCmd("holidays")
	cmd.NewFlag("state").
		String().
		Default("BY").
		Validator(func(v any) error {
			stateCodes := []string{"NATIONAL", "BW", "BY", "BE", "BB", "HB", "HH",
				"HE", "MV", "NI", "NW", "RP", "SL", "SN", "ST", "SH", "TH"}
			if v, ok := v.(string); !ok || !slices.Contains(stateCodes, v) {
				return errors.New("invalid state code")
			}
			return nil
		}).
		Help("a valid two-letter german state code")
	cmd.NewFlag("year").Int().Default(time.Now().Year()).Help("the year to fetch")

	// flax.F[string]("hall", nil, "some flag")
	// flag.FParse()
	// flax.F(flax.Foo[string]{"hall", "some flag"})

	if err := flax.Parse(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	// flag := cmd.Flag("state").ValueOr("state not set")
	// flag;cmd.Flag[int].Value
	flag := cmd.Flag("state").Value()
	fmt.Printf("parsed args: %#+v \n", flag)

	//    flax.SetFlag(flax.StringFlag{Name: "foo", Value: "hallo"})
	//        .Validator(func () {})
	// fooVal := flax.String("foo")
}
