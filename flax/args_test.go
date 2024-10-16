package flax

import (
	"fmt"
	"math/rand"
	"testing"
)

var argsInput = []string{"foobar", "bar", "test", "first", "nazz", "2323ff",
	"2323", "foobar-2323", "~/some/path/with/**/globbing", "long-arg-with-dashes"}

func shuffleArgsInput() {
	rand.Shuffle(len(argsInput), func(i int, j int) {
		tempI := argsInput[i]
		argsInput[i] = argsInput[j]
		argsInput[j] = tempI
	})
}

func TestShouldParsePositionalArgs(t *testing.T) {
	shuffleArgsInput()
	for i, in := range argsInput {
		raw := in
		arg := parseArg(raw, uint(i))
		if arg.Type == ArgTypeShort || arg.Type == ArgTypeLong || arg.Raw != raw || arg.Name != in {
			t.Fatalf("failed to parse short arg expected:%s:argc%d got:%v", raw, i, arg)
		}
	}
}

func TestShouldParseShortArgs(t *testing.T) {
	shuffleArgsInput()
	for i, in := range argsInput {
		raw := fmt.Sprintf("-%s", in)
		arg := parseArg(raw, uint(i))
		if arg.Type != ArgTypeShort || arg.Raw != raw || arg.Name != in {
			t.Fatalf("failed to parse short arg expected:%s:argc%d got:%v", raw, i, arg)
		}
	}
}

func TestShouldParseLongArgs(t *testing.T) {
	shuffleArgsInput()
	for i, in := range argsInput {
		raw := fmt.Sprintf("--%s", in)
		arg := parseArg(raw, uint(i))
		if arg.Type != ArgTypeLong || arg.Raw != raw || arg.Name != in {
			t.Fatalf("failed to parse long arg expected:%s:argc%d got:%v", raw, i, arg)
		}
	}
}

func TestShouldParseDounbleDash(t *testing.T) {
	raw := "--"
	arg := parseArg(raw, 1)
	if arg.Type != ArgTypeDoubleDash || arg.Raw != raw || arg.Name != "" {
		t.Fatalf("failed to parse double-dash arg expected: %s got: %v", raw, arg)
	}
}

func TestShouldParseProgram(t *testing.T) {
	raw := "foobar"
	arg := parseArg(raw, 0)
	if arg.Type != ArgTypeProgram || arg.Raw != raw || arg.Name != raw {
		t.Fatalf("failed to parse program arg expected: %s got: %v", raw, arg)
	}
}
