package flax

import "fmt"

func unreachable(cond bool, m string) {
	if cond {
		panic(fmt.Sprintf("UNREACHABLE: %s", m))
	}
}

func todo(m string) { panic(fmt.Sprintf("TODO: %s", m)) }
