package ies

import "rrc/utils"

// GapConfig-mgl-r16 ::= ENUMERATED
type GapconfigMglR16 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigMglR16EnumeratedNothing = iota
	GapconfigMglR16EnumeratedMs10
	GapconfigMglR16EnumeratedMs20
)
