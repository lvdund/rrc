package ies

// SPS-ConfigMulticastToReleaseList-r17 ::= SEQUENCE OF SPS-ConfigIndex-r16
// SIZE (1..8)
type SpsConfigmulticasttoreleaselistR17 struct {
	Value []SpsConfigindexR16 `lb:1,ub:8`
}
