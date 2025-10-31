package ies

// SRS-ResourceListConfigCLI-r16 ::= SEQUENCE OF SRS-ResourceConfigCLI-r16
// SIZE (1.. maxNrofCLI-SRS-Resources-r16)
type SrsResourcelistconfigcliR16 struct {
	Value []SrsResourceconfigcliR16 `lb:1,ub:maxNrofCLISrsResourcesR16`
}
