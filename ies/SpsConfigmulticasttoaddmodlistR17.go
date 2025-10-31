package ies

// SPS-ConfigMulticastToAddModList-r17 ::= SEQUENCE OF SPS-Config
// SIZE (1..8)
type SpsConfigmulticasttoaddmodlistR17 struct {
	Value []SpsConfig `lb:1,ub:8`
}
