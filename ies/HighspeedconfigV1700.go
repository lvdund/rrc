package ies

import "rrc/utils"

// HighSpeedConfig-v1700 ::= SEQUENCE
// Extensible
type HighspeedconfigV1700 struct {
	HighspeedmeascaScellR17   *utils.BOOLEAN
	HighspeedmeasinterfreqR17 *utils.BOOLEAN
	HighspeeddemodcaScellR17  *utils.BOOLEAN
}
