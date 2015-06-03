package manyflags

import (
	"os"
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestOverwriteArgs(t *testing.T) {
	expect := []string{"command", "-a", "-b", "-c"}
	os.Args = []string{"command", "-abc"}
	OverwriteArgs()
	if equal(expect, os.Args) {
		return
	}
	t.Errorf("expected: %v actual: %v", expect, os.Args)
}

func TestRemark(t *testing.T) {
	expect := []string{"-a", "-b", "-c", "10", "--", "-def"}
	arg := []string{"-a", "-bc", "10", "--", "-def"}
	actual := remake(arg)
	if equal(expect, actual) {
		return
	}
	t.Errorf("expected: %v actual: %v", expect, actual)
}

func TestSplitChunk(t *testing.T) {
	expect := []string{"-a", "-b", "-c"}
	actual := splitChunk("abc")
	if equal(expect, actual) {
		return
	}
	t.Errorf("expected: %v actual: %v", expect, actual)
}

func TestIsNormalFlag(t *testing.T) {
	expects := []bool{true, true, false, false}
	args := []string{"-x", "--x", "---x", ""}
	for i := range expects {
		actual := isNormalFlag(args[i])
		if expects[i] != actual {
			t.Errorf("expected: %v actual: %v", expects[i], actual)
		}
	}
}

func TestIsChunkFlag(t *testing.T) {
	expects := []bool{true, true, false, false, false}
	args := []string{"-abc", "--abc", "-8", "-a", "abc"}
	for i := range expects {
		actual := isChunkFlag(args[i])
		if expects[i] != actual {
			t.Errorf("expected: %v actual: %v", expects[i], actual)
		}
	}
}
