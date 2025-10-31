package ies

import "rrc/utils"

// MeasResultCDMA2000-measResult ::= SEQUENCE
// Extensible
type Measresultcdma2000Measresult struct {
	Pilotpnphase  *utils.INTEGER `lb:0,ub:32767`
	Pilotstrength utils.INTEGER  `lb:0,ub:63`
}
