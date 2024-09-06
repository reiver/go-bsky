package record

import (
	"github.com/reiver/go-bsky/record/registry"
)

func New(name string) (Record, bool) {
	var fn func()Record
	var found bool

	fn, found = registry.NewFuncs.Get(name)
	if !found {
		return nil, false
	}

	return fn(), true
}
