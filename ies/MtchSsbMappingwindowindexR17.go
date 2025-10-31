package ies

import "rrc/utils"

// MTCH-SSB-MappingWindowIndex-r17 ::= utils.INTEGER (0..maxNrofMTCH-SSB-MappingWindow-1-r17)
type MtchSsbMappingwindowindexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofMTCHSsbMappingwindow1R17`
}
