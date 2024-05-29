package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lukasjoc/flax/flax"
)

func isPrime(n int) string {
	if n <= 1 {
		return "No"
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	cli := flax.NewCLI("is-prime", "checks if the given number is prime")
	cli.Set(&flax.FlagSpec{
		Name:     "number",
		Short:    "n",
		Help:     "Some n in the range of 2..100",
		Required: true,
	})

	if err := cli.Parse(os.Args); err != nil {
		cli.Exit(err)
	}
	n, err := flax.IntFunc(
		func(v int) error {
			if v < 2 || v > 100 {
				return errors.New("invalid range for number flag: expected 2..100 (inclusive)")
			}
			return nil
		}).Try(cli.Get("number"))
	if err != nil {
		cli.Exit(err)
	}
	fmt.Printf("N: %d PRIME?: %s \n", n, isPrime(n))
}

// func main() {
// 	cli := flax.NewCLI("flax-tb", "simple cli for testing flax")
// 	cli.Set(&flax.FlagSpec{
// 		Name:     "firstname",
// 		Short:    "f",
// 		Required: true,
// 		Help:     "A given firstname of a certain predefined setup of names.",
// 	})
// 	cli.Set(&flax.FlagSpec{
// 		Name:     "lastname",
// 		Short:    "l",
// 		Required: true,
// 		Help:     "A given lastname of a certain predefined setup of lastnames.",
// 	})
// 	cli.Set(&flax.FlagSpec{
// 		Name:  "profession",
// 		Short: "p",
// 		Help:  "Provide your profession.",
// 	})
//
// 	if err := cli.Parse(os.Args); err != nil {
// 		cli.Exit(err)
// 	}
//
// 	// firstname, err := flax.String().Try(cli.Get("firstname"))
// 	// if err != nil {
// 	// 	cli.Exit(err)
// 	// }
// 	// fmt.Println("My firstname is: ", firstname)
//
// 	// lastname, err := flax.String().Try(cli.Get("lastname"))
// 	// if err != nil {
// 	// 	cli.Exit(err)
// 	// }
// 	// fmt.Println("My lastname is: ", lastname)
// 	//
// 	// profession, err := flax.StringFunc(func(v string) error {
// 	// 	if strings.Contains(v, "software") {
// 	// 		return errors.New("touch some grass")
// 	// 	}
// 	// 	return nil
// 	// }).Try(cli.Get("profession"))
// 	// if err != nil {
// 	// 	cli.Exit(err)
// 	// }
// 	// fmt.Println("My profession is: ", profession)
// }
