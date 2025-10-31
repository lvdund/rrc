package ies

import "rrc/utils"

// SRS-PosResourceId-r16 ::= utils.INTEGER (0..maxNrofSRS-PosResources-1-r16)
type SrsPosresourceidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRSPosresources1R16`
}
