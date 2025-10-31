package ies

// MTCH-SSB-MappingWindowList-r17 ::= SEQUENCE OF MTCH-SSB-MappingWindowCycleOffset-r17
// SIZE (1..maxNrofMTCH-SSB-MappingWindow-r17)
type MtchSsbMappingwindowlistR17 struct {
	Value []MtchSsbMappingwindowcycleoffsetR17 `lb:1,ub:maxNrofMTCHSsbMappingwindowR17`
}
