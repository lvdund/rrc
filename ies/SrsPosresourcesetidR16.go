package ies

import "rrc/utils"

// SRS-PosResourceSetId-r16 ::= utils.INTEGER (0..maxNrofSRS-PosResourceSets-1-r16)
type SrsPosresourcesetidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRSPosresourcesets1R16`
}
