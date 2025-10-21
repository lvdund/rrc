package ies

import "rrc/utils"

// AS-Context-NB ::= SEQUENCE
// Extensible
type AsContextNb struct {
	ReestablishmentinfoR13 *ReestablishmentinfoNb
}
