package record

import (
	"github.com/reiver/go-reg"
)

var NewFuncs reg.Registry[func()Record]

func New(name string) (Record, bool) {
	var fn func()Record
	var found bool

	fn, found = NewFuncs.Get(name)
	if !found {
		return nil, false
	}

	return fn(), true
}
