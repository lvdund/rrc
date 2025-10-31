package ies

// SL-ConfiguredGrantConfigList-r16 ::= SEQUENCE
type SlConfiguredgrantconfiglistR16 struct {
	SlConfiguredgrantconfigtoreleaselistR16 *[]SlConfigindexcgR16         `lb:1,ub:maxNrofCGSlR16`
	SlConfiguredgrantconfigtoaddmodlistR16  *[]SlConfiguredgrantconfigR16 `lb:1,ub:maxNrofCGSlR16`
}
