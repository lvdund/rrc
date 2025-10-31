package ies

import "rrc/utils"

// PH-UplinkCarrierMCG-ph-Type1or3 ::= ENUMERATED
type PhUplinkcarriermcgPhType1or3 struct {
	Value utils.ENUMERATED
}

const (
	PhUplinkcarriermcgPhType1or3EnumeratedNothing = iota
	PhUplinkcarriermcgPhType1or3EnumeratedType1
	PhUplinkcarriermcgPhType1or3EnumeratedType3
)
