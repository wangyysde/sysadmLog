// +build !appengine,!js,!windows,!nacl,!plan9

package sysadmLog

import (
	"io"
	"os"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}
