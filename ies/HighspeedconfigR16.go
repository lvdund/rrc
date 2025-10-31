package ies

import "rrc/utils"

// HighSpeedConfig-r16 ::= SEQUENCE
// Extensible
type HighspeedconfigR16 struct {
	HighspeedmeasflagR16  *utils.BOOLEAN
	HighspeeddemodflagR16 *utils.BOOLEAN
}
