package gomodtest

import (
	"github.com/rboyer/safeio"
)

func Remove(fn string) error {
	return safeio.Remove(fn)
}
