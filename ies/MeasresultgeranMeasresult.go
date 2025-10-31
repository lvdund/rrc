package ies

import "rrc/utils"

// MeasResultGERAN-measResult ::= SEQUENCE
// Extensible
type MeasresultgeranMeasresult struct {
	Rssi utils.INTEGER `lb:0,ub:63`
}
