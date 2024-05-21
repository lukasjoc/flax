package main

import (
	"errors"
	"slices"
	"time"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	// intflagscmd := flax.
	// 	NewCmd("basic").
	// 	    SetFlag("x", reflect.Int).Help("x for stuff").Cmd().
	// 	    SetFlag("y", reflect.Int).Help("y for stuff").Cmd().
	// 	NewSubCmd("geo").
	// 	    SetFlag("lat").Int().Help("lat for stuff").Cmd().
	// 	    SetFlag("long").Int().Help("long for stuff")
	flax.NewCmd("holidays").
		NewFlag("state").String().
		Default("BY").
		Validator(func(v any) error {
			stateCodes := []string{"NATIONAL", "BW", "BY", "BE", "BB", "HB", "HH",
				"HE", "MV", "NI", "NW", "RP", "SL", "SN", "ST", "SH", "TH"}
			if v, ok := v.(string); !ok || !slices.Contains(stateCodes, v) {
				return errors.New("invalid state code")
			}
			return nil
		}).
		Help("a valid two-letter german state code").Cmd().
		NewFlag("year").Int().
		Default(time.Now().Year()).
		Help("the year to fetch").Cmd().
		Parse()
	// flax.Parse(cmd)
	// fmt.Printf("cmd => %#+v \n", cmd)
}
