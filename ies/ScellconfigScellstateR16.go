package ies

import "rrc/utils"

// SCellConfig-sCellState-r16 ::= ENUMERATED
type ScellconfigScellstateR16 struct {
	Value utils.ENUMERATED
}

const (
	ScellconfigScellstateR16EnumeratedNothing = iota
	ScellconfigScellstateR16EnumeratedActivated
)
