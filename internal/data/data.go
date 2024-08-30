package data

import (
	"sort"
	"strings"
)

type T map[string]string

func (d T) String() string {
	ks := make([]string, 0, len(d))
	for k := range d {
		ks = append(ks, k)
	}

	sort.Strings(ks)

	return strings.Join(ks, "\n")
}
