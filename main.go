package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	cli := flax.NewCLI()
	cli.Set(&flax.FlagSpec{Name: "lat", Short: "la", Help: "define the lat of you location"})
	cli.Set(&flax.FlagSpec{Name: "long", Short: "lo", Help: "define the long of you location"})

	// parses the flags and puts the unparsed serial from the user to the thing
	// also fails if required flags were not set.. etc.. its doesnt do any
	// deserialization here
	cli.Parse(os.Args)

	posLat, err := flax.IntFunc(func(v int) error {
		if v < 0 {
			return errors.New("v is smaller than 0")
		}
		return nil
	}).Try(cli.Get("lat"))
	if err != nil {
		panic(err)
	}
	fmt.Println(posLat)

	// basic parsing to int
	lat, err := flax.Int().Try(cli.Get("lat"))
	if err != nil {
		panic(err)
	}
	fmt.Println(lat)

    // basic parsing to string (as this is just returning the serial)
	name, err := flax.String().Try(cli.Get("name"))
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// for cli.Parsed() {
	// 	flag := cli.Next()
	// 	switch flag.Name {
	// 	case "name":
	// 		fmt.Println("? name ?", flag)
	// 	case "lat":
	// 		fmt.Println("? lat ?", flag)
	// 	case "long":
	// 		fmt.Println("? long ?", flag)
	// 	}
	// }
}

// func main() {
// 	cmd := flax.NewCmd("holidays")
// 	cmd.NewFlag("state").
// 		String().
// 		Default("BY").
// 		Validator(func(v any) error {
// 			stateCodes := []string{"NATIONAL", "BW", "BY", "BE", "BB", "HB", "HH",
// 				"HE", "MV", "NI", "NW", "RP", "SL", "SN", "ST", "SH", "TH"}
// 			if v, ok := v.(string); !ok || !slices.Contains(stateCodes, v) {
// 				return errors.New("invalid state code")
// 			}
// 			return nil
// 		}).
// 		Help("a valid two-letter german state code")
// 	cmd.NewFlag("year").Int().Default(time.Now().Year()).Help("the year to fetch")
//
// 	// flax.F[string]("hall", nil, "some flag")
// 	// flag.FParse()
// 	// flax.F(flax.Foo[string]{"hall", "some flag"})
//
// 	if err := flax.Parse(cmd); err != nil {
// 		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
// 		os.Exit(1)
// 	}
// 	// flag := cmd.Flag("state").ValueOr("state not set")
// 	// flag;cmd.Flag[int].Value
// 	flag := cmd.Flag("state").Value()
//     fmt.Printf("parsed args: %#+v \n", flag)
//
// 	//    flax.SetFlag(flax.StringFlag{Name: "foo", Value: "hallo"})
// 	//        .Validator(func () {})
// 	// fooVal := flax.String("foo")
// }
