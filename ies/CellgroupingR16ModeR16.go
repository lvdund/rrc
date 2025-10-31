package ies

import "rrc/utils"

// CellGrouping-r16-mode-r16 ::= ENUMERATED
type CellgroupingR16ModeR16 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupingR16ModeR16EnumeratedNothing = iota
	CellgroupingR16ModeR16EnumeratedSync
	CellgroupingR16ModeR16EnumeratedAsync
)
