package ies

// MultiPDSCH-TDRA-r17 ::= SEQUENCE
// Extensible
type MultipdschTdraR17 struct {
	PdschTdraListR17 []PdschTimedomainresourceallocationR16 `lb:1,ub:maxNrofMultiplePDSCHsR17`
}
