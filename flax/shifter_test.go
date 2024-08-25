package flax

import (
	"strings"
	"testing"
)

func resetShifter() {
	shifter.n = 0
}

func init() {
	// TODO: Might want to convert to a constructor function instead, and pass as reference to
	// shifting functions. I really dont like init. :D
	args := strings.Fields("./foobar --bazz -bar")
	shifter = Shifter{args, uint(len(args)), 0}
}

func TestShouldShiftArg(t *testing.T) {
	resetShifter()
	arg := Shift()
	if arg == nil {
		t.Fatalf("failed shifting arg. expected arg but got: %v", arg)
	}
}

func TestShouldUnshiftArg(t *testing.T) {
	resetShifter()
	Shift() // shift away the program arg
	arg := Shift()
	if arg == nil {
		t.Fatalf("failed shifting arg. expected arg but got: %v", arg)
	}
	Unshift()
	arg1 := Shift()
	if arg1 == nil {
		t.Fatalf("failed shifting arg. expected arg but got: %v", arg1)
	}
	if arg1.Name != arg.Name {
		t.Fatalf("failed unshifting arg. expected arg:%s but got: %s", arg.Name, arg1.Name)
	}
}

func TestShouldPeekArg(t *testing.T) {
	resetShifter()
	Shift() // shift away the program arg
	arg := Peek()
	if arg == nil {
		t.Fatalf("failed shifting arg. expected arg but got: %v", arg)
	}
	arg1 := Shift()
	if arg1 == nil {
		t.Fatalf("failed shifting arg. expected arg but got: %v", arg1)
	}
	if arg1.Name != arg.Name {
		t.Fatalf("failed peeking arg. expected arg:%s but got: %s", arg.Name, arg1.Name)
	}
}
