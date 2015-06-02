package manyflags

import (
	"os"
	"regexp"
)

var flagRe = regexp.MustCompile(`^--?([a-zA-Z]{2,})$`)

// OverwriteArgs overwrite os.Args. It splits flags (e.g. -abc), and replace
// original argument. This function does not validate flag definition.
func OverwriteArgs() {
	var newArgs []string
	newArgs = append(newArgs, os.Args[0])
	newArgs = append(newArgs, remake(os.Args[1:])...)
	os.Args = newArgs
}

func remake(args []string) []string {
	var result []string
	for i, v := range args {
		// terminator. stop parsing.
		if v == "--" {
			result = append(result, args[i:]...)
			break
		}

		// just an argument or normal flag. ignore.
		if !isChunkFlag(v) || isNormalFlag(v) {
			result = append(result, args[i])
			continue
		}

		chunk := flagRe.FindStringSubmatch(v)
		result = append(result, splitChunk(chunk[1])...)
	}

	return result
}

func splitChunk(chunk string) []string {
	var flags []string
	for _, c := range chunk {
		flags = append(flags, "-"+string(c))
	}
	return flags
}

func isNormalFlag(v string) bool {
	// -x
	if len(v) == 2 && v[0] == '-' {
		return true
	}
	// --x
	if len(v) == 3 && v[0:2] == "--" {
		return true
	}
	return false
}

func isChunkFlag(v string) bool {
	return flagRe.MatchString(v)
}
