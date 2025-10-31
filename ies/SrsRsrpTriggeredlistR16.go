package ies

// SRS-RSRP-TriggeredList-r16 ::= SEQUENCE OF SRS-ResourceId
// SIZE (1.. maxNrofCLI-SRS-Resources-r16)
type SrsRsrpTriggeredlistR16 struct {
	Value []SrsResourceid `lb:1,ub:maxNrofCLISrsResourcesR16`
}
