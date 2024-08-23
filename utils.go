package biscuit

import "strings"

func cutAnyPrefix(s string, prefixes ...string) string {
	for _, p := range prefixes {
		s, _ = strings.CutPrefix(s, p)
	}

	return s
}