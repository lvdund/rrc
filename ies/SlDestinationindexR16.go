package ies

import "rrc/utils"

// SL-DestinationIndex-r16 ::= utils.INTEGER (0..maxNrofSL-Dest-1-r16)
type SlDestinationindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLDest1R16`
}
