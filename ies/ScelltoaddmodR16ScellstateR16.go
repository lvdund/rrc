package ies

import "rrc/utils"

// SCellToAddMod-r16-sCellState-r16 ::= ENUMERATED
type ScelltoaddmodR16ScellstateR16 struct {
	Value utils.ENUMERATED
}

const (
	ScelltoaddmodR16ScellstateR16EnumeratedNothing = iota
	ScelltoaddmodR16ScellstateR16EnumeratedActivated
	ScelltoaddmodR16ScellstateR16EnumeratedDormant
)
