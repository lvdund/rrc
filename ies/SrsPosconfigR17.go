package ies

// SRS-PosConfig-r17 ::= SEQUENCE
type SrsPosconfigR17 struct {
	SrsPosresourcesettoreleaselistR17 *[]SrsPosresourcesetidR16 `lb:1,ub:maxNrofSRSPosresourcesetsR16`
	SrsPosresourcesettoaddmodlistR17  *[]SrsPosresourcesetR16   `lb:1,ub:maxNrofSRSPosresourcesetsR16`
	SrsPosresourcetoreleaselistR17    *[]SrsPosresourceidR16    `lb:1,ub:maxNrofSRSPosresourcesR16`
	SrsPosresourcetoaddmodlistR17     *[]SrsPosresourceR16      `lb:1,ub:maxNrofSRSPosresourcesR16`
}
