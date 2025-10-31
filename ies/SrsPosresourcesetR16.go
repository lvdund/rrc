package ies

import "rrc/utils"

// SRS-PosResourceSet-r16 ::= SEQUENCE
// Extensible
type SrsPosresourcesetR16 struct {
	SrsPosresourcesetidR16    SrsPosresourcesetidR16
	SrsPosresourceidlistR16   *[]SrsPosresourceidR16 `lb:1,ub:maxNrofSRSResourcesperset`
	ResourcetypeR16           *SrsPosresourcesetR16ResourcetypeR16
	AlphaR16                  *Alpha
	P0R16                     *utils.INTEGER `lb:0,ub:24`
	PathlossreferencersPosR16 *SrsPosresourcesetR16PathlossreferencersPosR16
}
