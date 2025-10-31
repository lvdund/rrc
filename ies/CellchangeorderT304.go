package ies

import "rrc/utils"

// CellChangeOrder-t304 ::= ENUMERATED
type CellchangeorderT304 struct {
	Value utils.ENUMERATED
}

const (
	CellchangeorderT304EnumeratedNothing = iota
	CellchangeorderT304EnumeratedMs100
	CellchangeorderT304EnumeratedMs200
	CellchangeorderT304EnumeratedMs500
	CellchangeorderT304EnumeratedMs1000
	CellchangeorderT304EnumeratedMs2000
	CellchangeorderT304EnumeratedMs4000
	CellchangeorderT304EnumeratedMs8000
	CellchangeorderT304EnumeratedMs10000_V1310
)
